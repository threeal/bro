package cli

import (
	"context"
	"strings"

	"github.com/threeal/threeal-bot/pkg/schema"
	"google.golang.org/grpc"
)

type EchoClient struct {
	schema.EchoClient
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return EchoClient{schema.NewEchoClient(cc)}
}

func (c *EchoClient) Call(ctx context.Context, args []string) (string, error) {
	msg := strings.Join(args, " ")
	res, err := c.Echo(ctx, &schema.Message{Message: msg})
	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}
