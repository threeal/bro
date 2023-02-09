package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/threeal/bro/pkg/systemctl"
)

func getStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Status command",
		Long:  `A command to check Bro backend status..`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			info, err := systemctl.Status("bro")
			if err != nil {
				log.Fatalf("Failed to get status of the Bro Backend service: %v", err)
			}
			fmt.Printf("Status: %s\n", info.Active)
			fmt.Print(info.MainPid)
		},
	}
}
