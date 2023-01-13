package app

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/santhoshreddy97/goreddy/internal/errors"
)

func CreateDirectory(appName string) {
	_, err := os.Stat(appName)
	if err == nil {
		errors.CustomError("Directory by this app name is already present", err)
	}
	os.Mkdir(appName, DIRECTORY_PERMISSIONS)
}

// CreateGoFile creates the go file and writes the package name in it.
func CreateGoFile(packageName string, FileName string) {
	FileNameWithExt := fmt.Sprintf("%s.go", FileName)
	go_file_content := fmt.Sprintf("package %s\n", packageName)
	go_file_data := []byte(go_file_content)
	err := ioutil.WriteFile(packageName+"/"+FileNameWithExt, go_file_data, 0644)
	if err != nil {
		errors.CustomError("Error on creating file package", packageName, FileName)
	}
}
