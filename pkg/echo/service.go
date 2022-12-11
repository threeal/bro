package echo

import (
	"context"
)

type Server struct {
	UnimplementedEchoServer
}

func (s *Server) Echo(ctx context.Context, msg *Message) (*Message, error) {
	return msg, nil
}
