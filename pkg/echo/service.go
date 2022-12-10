package echo

import (
	"context"
	"strings"

	"google.golang.org/grpc"
)

type Server struct {
	UnimplementedEchoServer
}

func (s *Server) Echo(ctx context.Context, msg *Message) (*Message, error) {
	return msg, nil
}

func CallService(cc grpc.ClientConnInterface, ctx context.Context, args []string) (string, error) {
	c := NewEchoClient(cc)
	msg := strings.Join(args, " ")
	res, err := c.Echo(ctx, &Message{Message: msg})
	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}
