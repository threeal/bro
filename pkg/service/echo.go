package service

import (
	"context"

	"github.com/threeal/threeal-bot/pkg/echo"
)

type EchoServer struct {
	echo.UnimplementedEchoServer
}

func (s *EchoServer) Echo(ctx context.Context, msg *echo.Message) (*echo.Message, error) {
	return msg, nil
}
