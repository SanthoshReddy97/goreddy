package system

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func StartHttpServer(cmd *cobra.Command, args []string) {
	server := exec.Command("air")
	Execute(server, "groot")
}
