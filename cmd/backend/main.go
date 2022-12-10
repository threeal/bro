package main

import (
	"log"
	"net"
	"os"

	"threeal/threeal-bot/pkg/echo"

	"google.golang.org/grpc"
)

func getAddr() string {
	port, ok := os.LookupEnv("THREEAL_BOT_ADDR")
	if !ok {
		return ":50051"
	}
	return port
}

func main() {
	lis, err := net.Listen("tcp", getAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	echo.RegisterService(server)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
