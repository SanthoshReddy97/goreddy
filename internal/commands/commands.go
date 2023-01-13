package commands

import (
	"fmt"

	"github.com/santhoshreddy97/goreddy/internal/app"
	"github.com/spf13/cobra"
)

const (
	createApp = "createApp"
	version   = "version"
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
