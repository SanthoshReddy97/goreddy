package app

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/SanthoshReddy97/goreddy/internal/errors"
	"github.com/SanthoshReddy97/goreddy/internal/system"
	"github.com/SanthoshReddy97/goreddy/internal/templates"
)

// InitMod run the go mod init method.
// Running this go mod init command only to add the go version dynamically in go.mod file.
//
// Moving the go.mod file inside the appDirectory
func InitMod(appName string) {
	moduleName := fmt.Sprintf("github.com/gitUsername/%s", appName)
	cmda := exec.Command("go", "mod", "init", moduleName)
	system.Execute(cmda, appName)
}

// InitMain Creates the main.go file with basic information to start a server.
func InitMain(appName string) {
	err := ioutil.WriteFile(fmt.Sprintf("%s/manage.go", appName), templates.Main(), 0644)
	if err != nil {
		errors.CustomError("Error on creating main file", err)
	}
}

func InstallMainPackages(appName string) {
	gin := exec.Command("go", "get", "-u", "github.com/SanthoshReddy97/goreddy")
	system.Execute(gin, appName)
}

func InitAppSettings(appName string) {
	// creates the directory with appName
	cmd := exec.Command("mkdir", appName)
	system.Execute(cmd, appName)

	// Create settings file
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s/settings.go", appName, appName), templates.Settings(appName), 0644)
	if err != nil {
		errors.CustomError("Error on creating settings file", err)
	}

	// Create urls files
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/urls.go", appName, appName), templates.Urls(appName), 0644)
	if err != nil {
		errors.CustomError("Error on creating settings file", err)
	}
}
