package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goreddy",
	Short: "Go Rest Framework",
	Long:  `Go Rest Framework Tool`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createAppCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startServerCmd)
}
