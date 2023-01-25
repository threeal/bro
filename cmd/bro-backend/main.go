package main

import (
	"log"

	"github.com/threeal/bro/pkg/schema"
	"github.com/threeal/bro/pkg/service"
	"github.com/threeal/bro/pkg/tcp"
	"github.com/threeal/bro/pkg/utils"
)

func main() {
	backendConfig := utils.InitializeBackendConfig()
	addr := *backendConfig.ListenAddr
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
