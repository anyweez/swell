package main

import (
	"errors"
	"fmt"
	"github.com/anyweez/swell/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
)

type SwellProfile struct {
	Name          string
	AbsScriptPath []string
	Enabled       bool
}

/**
 * Tries to load the requested profiles from the location specified in the configuration.
 * The scripts are queued up alphabetically and executed when SwellProfile.Apply() is
 * called.
 */
func LoadProfiles(profileNames []string, config SwellConfig) ([]SwellProfile, error) {
	profiles := make([]SwellProfile, 0)

	potentials, err := ioutil.ReadDir(config.ProfilesPath)
	if err != nil {
		return profiles, err
	}

	for _, potential := range potentials {
		if potential.IsDir() {
			profile, err := loadProfile(potential.Name(), config)

			if err != nil {
				return profiles, errors.New(fmt.Sprintf("Error loading profile '%s': %s", profile.Name, err.Error()))
			}

			// Enable the profile if it was specified as one of the active profiles. Otherwise
			// leave it alone (it defaults to false).
			for _, name := range profileNames {
				if profile.Name == name {
					profile.Enabled = true
				}
			}

			profiles = append(profiles, profile)
		} else {
			logrus.Warn(potential.Name() + " is not a directory; skipping...")
		}
	}

	return profiles, nil
}

/**
 * Load a single profile. This is designed to be an internal function used by the public
 * LoadProfiles but is what does the heavy lifting for profile loading (LoadProfiles
 * simply invokes this once per profile).
 */
func loadProfile(name string, config SwellConfig) (SwellProfile, error) {
	fullDir := fmt.Sprintf("%s/%s", config.ProfilesPath, name)
	profile := SwellProfile{
		Name:    name,
		Enabled: false, // will be flipped to true later if it needs to be
	}

	src, err := os.Stat(fullDir)
	if err != nil {
		return SwellProfile{}, errors.New(fmt.Sprintf("Profile '%s' does not exist.", name))
	}

	if !src.IsDir() {
		return SwellProfile{}, errors.New(fmt.Sprintf("Profile doesn't have a corresponding directory."))
	}

	files, err := ioutil.ReadDir(fullDir)

	for _, file := range files {
		profile.AbsScriptPath = append(profile.AbsScriptPath, fmt.Sprintf("%s/%s", fullDir, file.Name()))
	}

	return profile, nil
}

/**
 * Execute all of the scripts associated with this profile. Apply() currently executes
 * all scripts serially, although it may become at least partially parallelizable in
 * the future.
 */
func (p *SwellProfile) Apply() error {
	for _, script := range p.AbsScriptPath {
		logrus.WithFields(logrus.Fields{
			"profile": p.Name,
			"script":  script,
		}).Info("Starting new task")

		cmd := exec.Command(script)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"profile": p.Name,
				"script":  script,
			}).Warn(fmt.Sprintf("Task failed: %s", err.Error()))
		} else {
			logrus.WithFields(logrus.Fields{
				"profile": p.Name,
				"script":  script,
			}).Info("Task complete")
		}
	}

	return nil
}
