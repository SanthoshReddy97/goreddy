package commands

import (
	"fmt"

	"github.com/SanthoshReddy97/goreddy/internal/system"
	"github.com/SanthoshReddy97/goreddy/src/app"
	"github.com/spf13/cobra"
)

const (
	createApp   = "createApp"
	version     = "version"
	startServer = "startServer"
)

var createAppCmd = &cobra.Command{
	Use:   createApp,
	Short: "Creates a blueprint of the app",
	Run:   app.CreateApp,
}

var versionCmd = &cobra.Command{
	Use:   version,
	Short: "Version of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 1.0.0")
	},
}

var startServerCmd = &cobra.Command{
	Use:   startServer,
	Short: "Starts the server",
	Run:   system.StartHttpServer,
}
