package service

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
	"github.com/threeal/bro/pkg/schema"
	"github.com/threeal/bro/pkg/tcp"
)

func TestEchoServer(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)
	server, err := tcp.NewServer(":" + strconv.Itoa(port))
	require.NoError(t, err)
	schema.RegisterEchoServer(server, &EchoServer{})
	go func() { server.Serve() }()
	defer server.Stop()
	time.Sleep(30 * time.Millisecond)
	conn, err := tcp.Connect("localhost:" + strconv.Itoa(port))
	require.NoError(t, err)
	defer conn.Close()
	client := schema.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	t.Run("CallEcho", func(t *testing.T) {
		msg := schema.Message{Message: "Hello world!"}
		res, err := client.Echo(ctx, &msg)
		require.NoError(t, err)
		require.Equal(t, msg.GetMessage(), res.GetMessage())
	})
}
