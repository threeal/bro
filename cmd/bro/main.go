package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/threeal/bro/pkg/bro"
	"github.com/threeal/bro/pkg/cli"
	"github.com/threeal/bro/pkg/tcp"
	"github.com/threeal/bro/pkg/utils"

	"google.golang.org/grpc"
)

func getClient(key string, conn grpc.ClientConnInterface) cli.Client {
	if key == "echo" {
		return cli.NewEchoClient(conn)
	}
	return nil
}

func main() {
	flag.Parse()
	config, err := utils.InitializeConfig(&bro.Config{})
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	addr := *config.(*bro.Config).BackendAddr
	conn, err := tcp.Connect(addr)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := getClient(flag.Arg(0), conn)
	if client == nil {
		log.Fatalf("invalid command: %s", flag.Arg(0))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Call(ctx, flag.Args()[1:])
	if err != nil {
		log.Fatalf("failed to call command: %v", err)
	}
	log.SetFlags(0)
	log.Println(res)
}
