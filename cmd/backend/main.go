package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "threeal/threeal-bot/pkg/echo"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	pb.UnimplementedEchoServer
}

func (s *Server) Echo(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	return msg, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listn: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterEchoServer(srv, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
