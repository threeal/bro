package tcp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	server, err := NewServer(":50050")
	require.NoError(t, err)
	go func() { server.Serve() }()
	defer server.Stop()
	conn, err := Connect("localhost:50050")
	require.NoError(t, err)
	require.NotNil(t, conn)
	err = conn.Close()
	require.NoError(t, err)
}
