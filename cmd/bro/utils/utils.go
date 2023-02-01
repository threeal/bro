package utils

import (
	"github.com/threeal/bro/pkg/cli"
	"google.golang.org/grpc"
)

func GetClient(key string, conn grpc.ClientConnInterface) cli.Client {
	if key == "echo" {
		return cli.NewEchoClient(conn)
	}
	return nil
}
