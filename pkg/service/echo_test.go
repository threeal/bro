package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/threeal/threeal-bot/pkg/schema"
	"github.com/threeal/threeal-bot/pkg/tcp"
)

func TestEchoServer(t *testing.T) {
	server, err := tcp.NewServer(":50050")
	require.NoError(t, err)
	schema.RegisterEchoServer(server, &EchoServer{})
	go func() { server.Serve() }()
	defer server.Stop()
	time.Sleep(30 * time.Millisecond)
	conn, err := tcp.Connect("localhost:50050")
	require.NoError(t, err)
	defer conn.Close()
	client := schema.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	t.Run("CallEcho", func(t *testing.T) {
		msg := schema.Message{Message: "Hello world!"}
		res, err := client.Echo(ctx, &msg)
		if err != nil {
			t.Fatalf("failed to call echo rpc: %v", err)
		}
		if res.GetMessage() != msg.GetMessage() {
			t.Fatalf("expected '%s', got: %s", msg.GetMessage(), res.GetMessage())
		}
	})
}
