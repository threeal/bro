package service

import (
	"context"

	"github.com/threeal/threeal-bot/pkg/schema"
)

type EchoServer struct {
	schema.UnimplementedEchoServer
}

func (s *EchoServer) Echo(ctx context.Context, msg *schema.Message) (*schema.Message, error) {
	return msg, nil
}
