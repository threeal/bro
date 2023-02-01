package commands

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/threeal/bro/cmd/bro/config"
	broUtils "github.com/threeal/bro/cmd/bro/utils"
	"github.com/threeal/bro/pkg/tcp"
	"github.com/threeal/bro/pkg/utils"
)

func getEchoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "echo [MESSAGE]",
		Short: "Bro",
		Long:  `Your friendly, personal, multi-purpose buddy written in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			msg := args			
			clientCmd := cmd.Name()
			config := config.Config{}
			err := utils.InitializeConfig(&config)
			if err != nil {
				log.Fatalf("failed to initialize config: %v", err)
			}
			addr := config.BackendAddr
			conn, err := tcp.Connect(addr)
			if err != nil {
				log.Fatalf("failed to connect: %v", err)
			}
			defer conn.Close()
			client := broUtils.GetClient(clientCmd, conn)
			if client == nil {
				log.Fatalf("invalid command: %s", clientCmd)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Call(ctx, msg)
			if err != nil {
				log.Fatalf("failed to call command: %v", err)
			}
			log.SetFlags(0)
			log.Println(res)
		},
	}
}
