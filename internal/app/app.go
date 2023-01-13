package app

import (
	"fmt"

	"github.com/santhoshreddy97/goreddy/internal/errors"
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
	InitMain(appName)
	log.Info(fmt.Sprintf("Hola! Created the app successfully!\n\ncd %s\ngo run main.go", appName))
}
