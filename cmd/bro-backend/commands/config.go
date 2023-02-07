package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/threeal/bro/cmd/bro-backend/config"
	"github.com/threeal/bro/pkg/utils"
)

func getConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Config command",
		Long:  `A utility command to configure bro backend..`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			conf := &config.Config{}
			if !utils.StringInSlice(key, utils.GetJSONFields(*conf)) {
				fmt.Printf("%s is not a valid argument", key)
				return
			}
			if len(args) < 2 {
				utils.ReadConfigFromFile(conf, config.CONFIG_FILENAME)
				fmt.Printf("%s: %s\n", key, conf.ListenAddr)
			} else {
				conf.ListenAddr = args[1]
				utils.WriteConfigToFile(conf, config.CONFIG_FILENAME)
				fmt.Printf("Successfully set %s to %s\n", key, conf.ListenAddr)
			}
		},
	}
}
