package brobackend

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	t.Run("It should init, write and read correctly when initialized with addr", func (t *testing.T) {
		addr := ":320"
		conf := Config{&addr}
		tmpDir := t.TempDir()
		t.Setenv("HOME", tmpDir)
		input := strings.NewReader(":69\n")
		err := conf.Init(input)
		require.NoError(t, err)
		err = conf.Write()
		require.NoError(t, err)
		err = conf.Read()
		require.NoError(t, err)
	})
	t.Run("It should init correctly when given address", func (t *testing.T) {
		conf := Config{}
		input := strings.NewReader(":69\n")
		err := conf.Init(input)
		require.NoError(t, err)
	})
	t.Run("It should init correctly when address is empty", func (t *testing.T) {
		conf := Config{}
		input := strings.NewReader("\n")
		err := conf.Init(input)
		require.NoError(t, err)
	})
	t.Run("It should error when input is not valid", func (t *testing.T) {
		conf := Config{}
		input := strings.NewReader("\t")
		err := conf.Init(input)
		require.Error(t, err)
	})
}
