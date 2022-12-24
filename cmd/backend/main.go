package main

import (
	"log"

	"github.com/threeal/threeal-bot/pkg/schema"
	"github.com/threeal/threeal-bot/pkg/service"
	"github.com/threeal/threeal-bot/pkg/tcp"
	"github.com/threeal/threeal-bot/pkg/utils"
)

func main() {
	addr := utils.GetEnvOrDefault("THREEAL_BOT_ADDR", ":50051")
	server, err := tcp.NewServer(addr)
	if err != nil {
		log.Fatalf("failed to create a new server on `%s`: %v", addr, err)
	}
	schema.RegisterEchoServer(server, &service.EchoServer{})
	log.Printf("server listening at %v", server.Addr())
	if err := server.Serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
