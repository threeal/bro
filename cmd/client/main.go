package main

import (
	"context"
	"flag"
	"log"
	"time"

	"threeal/threeal-bot/pkg/echo"
	"threeal/threeal-bot/pkg/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Cmd = func(grpc.ClientConnInterface, context.Context, []string) (string, error)

var (
	cmds = map[string]Cmd{
		"echo": echo.CallService,
	}
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	cmd, ok := cmds[arg]
	if !ok {
		log.Fatalf("invalid command: %s", arg)
	}
	addr := utils.GetEnvOrDefault("THREEAL_BOT_ADDR", "localhost:50051")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := cmd(conn, ctx, flag.Args()[1:])
	if err != nil {
		log.Fatalf("failed to call command: %v", err)
	}
	log.SetFlags(0)
	log.Println(res)
}
