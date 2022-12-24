package tcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("NewServer", func(t *testing.T) {
		server, err := NewServer(":50050")
		assert.NoError(t, err)
		assert.NotNil(t, server)
		t.Run("Serve", func(t *testing.T) {
			res := make(chan error)
			go func() { res <- server.Serve() }()
			server.GracefulStop()
			err = <-res
			assert.ErrorContains(t, err, "the server has been stopped")
		})
	})
	t.Run("NewServerFailure", func(t *testing.T) {
		_, err := NewServer("invalid")
		assert.Error(t, err)
	})
}
