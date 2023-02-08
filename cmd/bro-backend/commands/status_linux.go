package commands

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func getStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Status command",
		Long:  `A command to check Bro backend status..`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			broService := exec.Command("systemctl", "status", "bro")
			out, _ := broService.CombinedOutput()
			activeRegex, _ := regexp.Compile("Active:.*")
			activeString := activeRegex.FindString(string(out))
			if activeString == "" {
				log.Fatalf(string(out))
			}
			status := strings.Split(activeString, ": ")[1]
			fmt.Printf("Status: %s\n", status)
			PIDRegex, _ := regexp.Compile("Main PID:.*")
			mainPIDString := PIDRegex.FindString(string(out))
			if mainPIDString != "" {
				fmt.Println(mainPIDString)
			}
		},
	}
}
