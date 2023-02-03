package commands

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	broUtils "github.com/threeal/bro/cmd/bro/utils"
)

func getEchoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "echo [messages...]",
		Short: "Echo command",
		Long:  `A command that outputs the strings that are passed to it as arguments.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			msg := args
			clientCmd := cmd.Name()
			conn, err := broUtils.ConnectToBackend()
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
