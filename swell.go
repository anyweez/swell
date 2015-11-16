package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"os"
)

/**
* Swell should:
 *  X Read in configuration settings from swell.json, checking the current directory first and
 *    falling back to /etc/swell.conf.
 *  - Then check the command line params for the desired profiles.
 *  - For each profile, check to see that a directory exists at $CONFIG_BASE_PATH/$PROFILE and
 *    read all scripts.
 *  - Arrange the available scripts into an execution chain and start running them in forked processes.
 *  - Log all errors.
*/

var PROFILES_PATH = flag.String("profiles", "profiles/", "Root directory where all profiles are stored.")

func main() {
	flag.Parse()

	// Validate that this is an acceptable path. Throw an error if not.
	path, err := GetProfilesPath(*PROFILES_PATH)
	if err != nil {
		logrus.Fatal(err.Error())
	} else {
		logrus.Info("Reading profiles from " + path)
	}

	// Create a configuration object, which for the time being only contains this path.
	config := SwellConfig{
		ProfilesPath: path,
	}

	// Load profiles
	profiles, err := LoadProfiles(os.Args[1:], config)
	if err != nil {
		logrus.Fatal("Error reading profiles: " + err.Error())
	}

	logrus.Info("Available profiles:")
	for _, profile := range profiles {
		if profile.Enabled {
			logrus.Info("   [run ] " + profile.Name)
		} else {
			logrus.Warn("   [skip] " + profile.Name)
		}
	}

	logrus.Info("Running profiles:")
	// Apply profiles in serial, and run all scripts serially for now as well.
	for _, profile := range profiles {
		if profile.Enabled {
			logrus.Info("   [run ] " + profile.Name)
			err = profile.Apply()

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"profile": profile.Name,
				}).Warn(err.Error())
			}
		} else {
			logrus.Warn("   [skip] " + profile.Name)
		}
	}

	logrus.Info("Complete!")
}
