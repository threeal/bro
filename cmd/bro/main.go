package main

import (
	"github.com/threeal/bro/pkg/cli"

	"google.golang.org/grpc"
)

func getClient(key string, conn grpc.ClientConnInterface) cli.Client {
	if key == "echo" {
		return cli.NewEchoClient(conn)
	}
	return nil
}

func main() {
	rootCmd.Execute()
}
