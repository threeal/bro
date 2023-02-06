package commands

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/threeal/bro/cmd/bro-backend/config"
	backendConfig "github.com/threeal/bro/cmd/bro-backend/config"
	"github.com/threeal/bro/pkg/utils"
)

func getStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Status command",
		Long:  `A command to check Bro backend if config exist and the service is running..`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if runtime.GOOS != "linux" {
				log.Fatalf("This command only available on Linux.")
			}
			config := &config.Config{}
			if err := utils.ReadConfigFromFile(config, backendConfig.CONFIG_FILENAME); err == nil {
				fmt.Printf("Bro config file found!, listen addr = %s\n", config.ListenAddr)
			} else {
				fmt.Println("Bro config file not found, consider running the spin command")
			}
			broService := exec.Command("systemctl", "check", "bro")
			out, _ := broService.CombinedOutput()
			fmt.Printf("Bro service status is: %s\n", string(out))
		},
	}
}
