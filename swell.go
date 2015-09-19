package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
)

/**
* Swell should:
 *  - Read in configuration settings from swell.conf, checking the current directory first and
 *    falling back to /etc/swell.conf.
 *  - Then check the command line params for the desired profiles.
 *  - For each profile, check to see that a directory exists at $CONFIG_BASE_PATH/$PROFILE and
 *    read all scripts.
 *  - Arrange the available scripts into an execution chain and start running them in forked processes.
 *  - Log all errors.
*/

func main() {
	// TODO: catch this error I suppose? Not sure when it would actually be thrown.
	local_dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	config, err := ReadConfig("swell.conf", []string{local_dir, "/etc"})
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Error reading configuration file: %s", err.Error()))
	}

	// Load profiles
	profiles, err := LoadProfiles(os.Args[1:], config)
	if err != nil {
		logrus.Fatal("Error reading profiles: " + err.Error())
	}

	// Apply profiles in serial, and run all scripts serially for now as well.
	for _, profile := range profiles {
		err = profile.Apply()

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"profile": profile.Name,
			}).Warn(err.Error())
		}
	}
}
