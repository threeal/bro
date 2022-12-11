package cli

import "context"

type Client interface {
	Call(ctx context.Context, args []string) (string, error)
}
