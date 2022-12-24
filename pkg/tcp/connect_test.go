package tcp

import (
	"strconv"
	"testing"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)
	server, err := NewServer(":" + strconv.Itoa(port))
	require.NoError(t, err)
	go func() { server.Serve() }()
	defer server.Stop()
	conn, err := Connect("localhost:" + strconv.Itoa(port))
	require.NoError(t, err)
	require.NotNil(t, conn)
	err = conn.Close()
	require.NoError(t, err)
}
