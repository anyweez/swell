package main

type SwellConfig struct {
    // The absolute path to the parent directory containing all profiles.
    ProfilesPath    string
}

/**
 * Checks a few different locations for a configuration file and parses
 * it if it finds it, returning a SwellConfig object.
 */
func ReadConfig(filename string, preferredPaths []string) (SwellConfig, error) {
    config := SwellConfig{
        ProfilesPath: "/Users/luke/git/swell/profiles",
    }
    
    return config, nil
}