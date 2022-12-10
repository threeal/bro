package echo

import (
	"context"

	"google.golang.org/grpc"
)

type Server struct {
	UnimplementedEchoServer
}

func (s *Server) Echo(ctx context.Context, msg *Message) (*Message, error) {
	return msg, nil
}

func RegisterService(s grpc.ServiceRegistrar) {
	RegisterEchoServer(s, &Server{})
}
