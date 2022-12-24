package tcp

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	net.Listener
}

func (s *Server) Serve() error {
	return s.Server.Serve(s.Listener)
}

func NewServer(address string) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	server := &Server{grpc.NewServer(), lis}
	return server, nil
}
