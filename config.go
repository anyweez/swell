package main

import (
	"errors"
	"os"
	"path/filepath"
)

type SwellConfig struct {
	// The absolute path to the parent directory containing all profiles.
	ProfilesPath string
}

/**
 * The user provides a directory that should be the root directory for all
 * usable profiles. This function confirms that the directory exists and
 * returns the absolute path. Users can pass in absolute or relative
 * paths.
 */
func GetProfilesPath(explicit string) (string, error) {
	provided, err := filepath.Abs(explicit)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(provided)
	// If there's an error stat'ing the FS node (i.e. doesn't exist), return the error.
	if err != nil {
		return "", err
		// Victory condition.
	} else if info.IsDir() {
		return provided, nil
		// If not a directory, we can't use it.
	} else {
		return "", errors.New("Specified profiles path is a file; must be a directory.")
	}
}
