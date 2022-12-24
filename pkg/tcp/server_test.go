package tcp

import (
	"strconv"
	"testing"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	t.Run("NewServer", func(t *testing.T) {
		port, err := freeport.GetFreePort()
		require.NoError(t, err)
		server, err := NewServer(":" + strconv.Itoa(port))
		require.NoError(t, err)
		require.NotNil(t, server)
		t.Run("Serve", func(t *testing.T) {
			res := make(chan error)
			go func() { res <- server.Serve() }()
			server.GracefulStop()
			err = <-res
			require.ErrorContains(t, err, "the server has been stopped")
		})
	})
	t.Run("NewServerFailure", func(t *testing.T) {
		_, err := NewServer("invalid")
		require.Error(t, err)
	})
}
