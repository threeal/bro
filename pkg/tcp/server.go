package tcp

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Lis net.Listener
	Srv *grpc.Server
}

func (s *Server) Serve() error {
	return s.Srv.Serve(s.Lis)
}

func NewServer(address string) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	server := &Server{Lis: lis, Srv: grpc.NewServer()}
	return server, nil
}
