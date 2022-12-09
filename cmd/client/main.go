package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	pb "threeal/threeal-bot/pkg/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	msg := strings.Join(flag.Args()[:], " ")
	r, err := c.Echo(ctx, &pb.Message{Message: msg})
	if err != nil {
		log.Fatalf("could not call echo: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
