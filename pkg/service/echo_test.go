package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/threeal/threeal-bot/pkg/schema"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestEchoServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		t.Fatalf("failed to listen to ':50050': %v", err)
	}
	server := grpc.NewServer()
	schema.RegisterEchoServer(server, &EchoServer{})
	go func() { server.Serve(lis) }()
	time.Sleep(100 * time.Millisecond)
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
