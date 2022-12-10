package main

import (
	"log"
	"net"

	"github.com/threeal/threeal-bot/pkg/echo"
	"github.com/threeal/threeal-bot/pkg/utils"

	"google.golang.org/grpc"
)

func main() {
	addr := utils.GetEnvOrDefault("THREEAL_BOT_ADDR", ":50051")
	lis, err := net.Listen("tcp", addr)
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
