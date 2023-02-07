package commands

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func getStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Status command",
		Long:  `A command to check Bro backend if config exist and the service is running..`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			broService := exec.Command("systemctl", "check", "bro")
			out, _ := broService.CombinedOutput()
			fmt.Printf("Bro service status is: %s\n", string(out))
		},
	}
}
