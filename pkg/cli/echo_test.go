package cli

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
	"github.com/threeal/threeal-bot/pkg/schema"
	"github.com/threeal/threeal-bot/pkg/service"
	"github.com/threeal/threeal-bot/pkg/tcp"
)

func TestEchoClient(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)
	server, err := tcp.NewServer(":" + strconv.Itoa(port))
	require.NoError(t, err)
	schema.RegisterEchoServer(server, &service.EchoServer{})
	go func() { server.Serve() }()
	defer server.Stop()
	time.Sleep(30 * time.Millisecond)
	conn, err := tcp.Connect("localhost:" + strconv.Itoa(port))
	require.NoError(t, err)
	defer conn.Close()
	client := NewEchoClient(conn)
	t.Run("Call", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		msg := "Hello world!"
		res, err := client.Call(ctx, []string{msg})
		if err != nil {
			t.Fatalf("echo client failed to call: %v", err)
		}
		if res != msg {
			t.Fatalf("expected '%s', got: %s", msg, res)
		}
	})
	t.Run("InvalidCall", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
		defer cancel()
		_, err := client.Call(ctx, []string{"foo"})
		if err == nil {
			t.Fatalf("expected echo client failed to call, got success instead")
		}
	})
}
