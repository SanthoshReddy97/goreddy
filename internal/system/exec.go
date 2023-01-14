package system

import (
	"os/exec"

	"github.com/SanthoshReddy97/goreddy/internal/errors"
)

func Execute(cmd *exec.Cmd, dirName string) {
	cmd.Dir = dirName
	_, err := cmd.Output()
	if err != nil {
		errors.CustomError("Error on executing command", cmd.Stdout, err)
	}
}
