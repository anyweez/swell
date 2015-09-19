package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type SwellConfig struct {
	// The absolute path to the parent directory containing all profiles.
	ProfilesPath string
}

/**
 * Checks a few different locations for a configuration file and parses
 * it if it finds it, returning a SwellConfig object.
 */
func ReadConfig(filename string, preferredPaths []string) (SwellConfig, error) {
	for _, path := range preferredPaths {
		fullPath := fmt.Sprintf("%s/%s", path, filename)
		fmt.Println(fullPath)
		_, err := os.Stat(fullPath)

		if err == nil {
			fp, err := ioutil.ReadFile(fullPath)

			// We know that the file exists, but it's possible we might hit
			// permission issues or something else.
			if err != nil {
				return SwellConfig{}, err
			}

			config := SwellConfig{}
			err = json.Unmarshal(fp, &config)

			if err != nil {
				return config, err
			}

			return config, nil
		}
	}

	return SwellConfig{}, errors.New(fmt.Sprintf("No valid %s file found", filename))
}
