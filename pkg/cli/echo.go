package cli

import (
	"context"
	"strings"

	"github.com/threeal/threeal-bot/pkg/schema"
	"google.golang.org/grpc"
)

type EchoClient struct {
	Client
	client schema.EchoClient
}

func NewEchoClient(cc grpc.ClientConnInterface) Client {
	return &EchoClient{client: schema.NewEchoClient(cc)}
}

func (c *EchoClient) Call(ctx context.Context, args []string) (string, error) {
	msg := strings.Join(args, " ")
	res, err := c.client.Echo(ctx, &schema.Message{Message: msg})
	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}
