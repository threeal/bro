package cli

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/threeal/threeal-bot/pkg/schema"
	"github.com/threeal/threeal-bot/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestEchoClient(t *testing.T) {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		t.Fatalf("failed to listen to ':50050': %v", err)
	}
	server := grpc.NewServer()
	schema.RegisterEchoServer(server, &service.EchoServer{})
	go func() { server.Serve(lis) }()
	time.Sleep(100 * time.Millisecond)
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect to 'localhost:50050': %v", err)
	}
	client := NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	msg := "Hello world!"
	t.Run("Call", func(t *testing.T) {
		res, err := client.Call(ctx, []string{msg})
		if err != nil {
			t.Fatalf("echo client failed to call: %v", err)
		}
		if res != msg {
			t.Fatalf("expected '%s', got: %s", msg, res)
		}
	})
}
