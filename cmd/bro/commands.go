package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/threeal/bro/pkg/tcp"
	"github.com/threeal/bro/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "echo [MESSAGE]",
	Short: "Bro",
	Long: `Your friendly, personal, multi-purpose buddy written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := args[1:]
		clientCmd := args[0]
		config := Config{}
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
		client := getClient(clientCmd, conn)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
