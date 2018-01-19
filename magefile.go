// +build mage

// nolint
package main

import (
	"fmt"
	"os"

	"github.com/aporeto-inc/domingo/golang"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var baseDir string

func init() {
	domingo.SetProjectName("elemental")

	bd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	baseDir = bd
}

// Init initialize the project.
func Init() {
	mg.Deps(
		domingo.Init,
	)
}

// Test runs unit tests.
func Test() {

	mg.Deps(
		domingo.Lint,
		domingo.Test,
	)
}

// ElegenBinarize inlines the elegen templates into a binary file.
func ElegenBinarize() error {

	if err := os.Chdir("cmd/elegen"); err != nil {
		return err
	}
	defer os.Chdir(baseDir)

	if err := os.MkdirAll("./static", 0755); err != nil {
		return err
	}

	if out, err := sh.Output("go-bindata", "-pkg", "static", "./templates"); err != nil {
		fmt.Println(out)
		return err
	}

	if err := os.RemoveAll("./static/bindata.go"); err != nil {
		return err
	}

	return sh.Run("mv", "./bindata.go", "./static/bindata.go")
}

func ElegenLinux() error { return elegenBuild("linux", domingo.BuildLinux) }

func ElegenDarwin() error { return elegenBuild("darwin", domingo.BuildDarwin) }

func Elegen() {
	mg.SerialDeps(
		ElegenBinarize,
		ElegenLinux,
		ElegenDarwin,
	)
}

func elegenBuild(platform string, buildFunc func() error) error {

	if err := os.Chdir("cmd/elegen"); err != nil {
		return err
	}
	defer os.Chdir(baseDir)

	if err := os.MkdirAll("./build/"+platform, 0755); err != nil {
		return err
	}

	if err := buildFunc(); err != nil {
		return err
	}

	return sh.Run("mv", "elegen", "build/"+platform+"/elegen")
}
