package main

import (
	"os"
	//    "path/filepath"
	"github.com/Sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var CURRENT_DIR, _ = os.Getwd()

func TestMain(t *testing.T) {
	// Create a local directory that matches the name of the default profiles
	// directory. This will be automatically deleted again by the gotest
	// executable when it clears all test files.
	//
	// Additionally, create a profile called "first."
	err := os.Mkdir(*PROFILES_PATH, os.ModeDir|0777)
	err = os.Mkdir(*PROFILES_PATH+"first", os.ModeDir|os.ModePerm)

	if err != nil {
		logrus.Info(os.Getwd())
		logrus.Error(err)
	}

	Convey("directories actually exists?", t, func() {
		workingDir, _ := os.Getwd()
		info, err := os.Stat(workingDir + "/" + *PROFILES_PATH)
		So(err, ShouldBeNil)
		So(info.IsDir(), ShouldEqual, true)

		info, err = os.Stat(workingDir + "/" + *PROFILES_PATH + "first")
		So(err, ShouldBeNil)
		So(info.IsDir(), ShouldEqual, true)
	})
}

/****
 ** config.go:GetConfigPath
 ****/
func TestGetProfilesPath(t *testing.T) {
	// When no special path is provided, this should return the default in
	// the current directory.
	// TODO: need this test but the testing directory mis-match is screwing it up atm
	/*
	   Convey("returns correct default path", t, func() {
	       path, err := GetProfilesPath(*PROFILES_PATH)

	       So(err, ShouldBeNil)
	       So(path, ShouldEqual, CURRENT_DIR + "/" + *PROFILES_PATH)
	   })
	*/

	/*
	   Convey("rejects paths that lead to files", t, func() {

	   })
	*/
	// When an absolute path is provided, don't append a bunch of garbage
	// in front of it for some reason. The output should be the same as
	// the input in this case (unless the file doesn't exist, then it should
	// return an error.
	Convey("doesn't mangle absolute directories", t, func() {
		path, err := GetProfilesPath(CURRENT_DIR)

		So(err, ShouldBeNil)
		So(path, ShouldEqual, CURRENT_DIR)
	})

	Convey("can detect existing absolute paths", t, func() {
		path, err := GetProfilesPath(CURRENT_DIR + "/profiles")

		So(err, ShouldBeNil)
		So(path, ShouldEqual, CURRENT_DIR+"/profiles")
	})

	Convey("returns an error when the directory doesn't exist", t, func() {
		path, err := GetProfilesPath("where-is-this")

		So(err, ShouldNotBeNil)
		So(path, ShouldEqual, "")
	})
}

func TestLoadProfiles(t *testing.T) {
	Convey("detects all profiles", t, func() {
		path, err := GetProfilesPath(CURRENT_DIR + "/profiles")

		profiles, err := LoadProfiles([]string{}, SwellConfig{
			ProfilesPath: path,
		})

		So(CURRENT_DIR+"/profiles", ShouldBeNil)
		So(err, ShouldBeNil)
		So(len(profiles), ShouldEqual, 1)
	})
}
