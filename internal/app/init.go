package app

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/santhoshreddy97/goreddy/internal/errors"
	"github.com/santhoshreddy97/goreddy/internal/templates"
)

// InitMod run the go mod init method.
// Running this go mod init command only to add the go version dynamically in go.mod file.
//
// Moving the go.mod file inside the appDirectory
func InitMod(appName string) {
	moduleName := fmt.Sprintf("github.com/gitUsername/%s", appName)
	cmda := exec.Command("go", "mod", "init", moduleName)
	_, err := cmda.Output()
	if err != nil {
		errors.CustomError("Error on initializing mod ", err)
	}
	mvCmd := exec.Command("mv", "go.mod", appName)
	_, err = mvCmd.Output()
	if err != nil {
		errors.CustomError("Error on moving go.mod", err)
	}
}

// InitMain Creates the main.go file with basic information to start a server.
func InitMain(appName string) {
	err := ioutil.WriteFile(fmt.Sprintf("%s/main.go", appName), templates.Main(), 0644)
	if err != nil {
		errors.CustomError("Error on creating main file", err)
	}
}
