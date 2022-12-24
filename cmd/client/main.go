package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/threeal/threeal-bot/pkg/cli"
	"github.com/threeal/threeal-bot/pkg/tcp"
	"github.com/threeal/threeal-bot/pkg/utils"

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
	addr := utils.GetEnvOrDefault("THREEAL_BOT_ADDR", "localhost:50051")
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
