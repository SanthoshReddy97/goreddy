package app

import (
	"fmt"

	"github.com/SanthoshReddy97/goreddy/internal/errors"
	"github.com/SanthoshReddy97/goreddy/src/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CreateApp starts creating the new go app.
// This is the startup method for the createApp shell command.
func CreateApp(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		errors.CustomError("Provide the app name", "")
	}
	appName := args[0]
	CreateDirectory(appName)
	InitMod(appName)
	InstallMainPackages(appName)
	InitMain(appName)
	log.Info(fmt.Sprintf("Hola! Created the app successfully!\n\ncd %s\ngo run main.go", appName))
}

func ExecuteFromCommandLine() {
	gin.StartHttpServer()
}
