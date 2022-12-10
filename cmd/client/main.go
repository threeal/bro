package main

import (
	"context"
	"flag"
	"log"
	"time"

	"threeal/threeal-bot/pkg/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	cmds = map[string]func(grpc.ClientConnInterface, context.Context, []string) (string, error){
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
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
