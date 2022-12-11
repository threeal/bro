package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/threeal/threeal-bot/pkg/cli"
	"github.com/threeal/threeal-bot/pkg/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	newClientFuncs = map[string]func(grpc.ClientConnInterface) cli.EchoClient{
		"echo": cli.NewEchoClient,
	}
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	newClientFunc, ok := newClientFuncs[arg]
	if !ok {
		log.Fatalf("invalid command: %s", arg)
	}
	addr := utils.GetEnvOrDefault("THREEAL_BOT_ADDR", "localhost:50051")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := newClientFunc(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Call(ctx, flag.Args()[1:])
	if err != nil {
		log.Fatalf("failed to call command: %v", err)
	}
	log.SetFlags(0)
	log.Println(res)
}
