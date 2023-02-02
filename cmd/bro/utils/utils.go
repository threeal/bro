package utils

import (
	"github.com/threeal/bro/cmd/bro/config"
	"github.com/threeal/bro/pkg/cli"
	"github.com/threeal/bro/pkg/tcp"
	"github.com/threeal/bro/pkg/utils"
	"google.golang.org/grpc"
)

func GetClient(key string, conn grpc.ClientConnInterface) cli.Client {
	if key == "echo" {
		return cli.NewEchoClient(conn)
	}
	return nil
}

func ConnectToBackend() (*grpc.ClientConn, error) {
	config := config.Config{}
	err := utils.InitializeConfig(&config)
	if err != nil {
		return nil, err
	}
	addr := config.BackendAddr
	return tcp.Connect(addr)
}
