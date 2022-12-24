package tcp

import (
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("NewServer", func(t *testing.T) {
		server, err := NewServer(":50050")
		if err != nil {
			t.Fatalf("failed to create a new server on `:50050`: %v", err)
		}
		if server == nil {
			t.Fatal("Expected `server` to be non-nil, got nil instead")
		}
		t.Run("Serve", func(t *testing.T) {
			res := make(chan error)
			go func() { res <- server.Serve() }()
			server.GracefulStop()
			err = <-res
			if !strings.Contains(err.Error(), "the server has been stopped") {
				t.Fatalf("failed during serving the server: %v", err)
			}
		})
	})
	t.Run("NewServerFailure", func(t *testing.T) {
		_, err := NewServer("invalid")
		if err == nil {
			t.Fatal("Expected `err` to be non-nil, got nil instead")
		}
	})
}
