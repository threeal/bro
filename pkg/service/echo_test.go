package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/threeal/threeal-bot/pkg/schema"
	"github.com/threeal/threeal-bot/pkg/tcp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestEchoServer(t *testing.T) {
	server, err := tcp.NewServer(":50050")
	assert.NoError(t, err)
	schema.RegisterEchoServer(server, &EchoServer{})
	go func() { server.Serve() }()
	time.Sleep(30 * time.Millisecond)
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect to 'localhost:50050': %v", err)
	}
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
