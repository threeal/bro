package tcp

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
	server   *grpc.Server
}

func (s *Server) Serve() error {
	return s.server.Serve(s.listener)
}

func NewServer(address string) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	server := &Server{listener: lis, server: grpc.NewServer()}
	return server, nil
}
