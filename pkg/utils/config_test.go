package utils

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testConfig struct {
	BackendAddr *string `json:"backend_addr"`
}

func (c *testConfig) Read() error {
	return nil
}

func (c *testConfig) Write() error {
	return nil
}

func (c *testConfig) Init(rd io.Reader) error {
	return nil
}

type errorConfig struct {
	BackendAddr *string `json:"backend_addr"`
}

func (c *errorConfig) Read() error {
	return errors.New("read error")
}

func (c *errorConfig) Write() error {
	return errors.New("write error")
}

func (c *errorConfig) Init(rd io.Reader) error {
	return errors.New("init error")
}

type funcFieldConfig struct {
	BackendAddr func() `json:"backend_addr"`
}

func (c *funcFieldConfig) Read() error {
	return errors.New("read error")
}

func (c *funcFieldConfig) Write() error {
	return errors.New("write error")
}

func (c *funcFieldConfig) Init(rd io.Reader) error {
	return errors.New("init error")
}

func TestConfig(t *testing.T) {
	filename := "ozymandias"
	addr := ":21"
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	t.Run("it should successfully create config flow", func (t *testing.T) {
		conf := &testConfig{&addr}
		_, err := InitializeConfig(conf)
		require.NoError(t, err)
		require.Equal(t, addr, *conf.BackendAddr)
		err = WriteConfigToFile(conf, filename)
		require.NoError(t, err)
		err = ReadConfigFromFile(conf, filename)
		require.NoError(t, err)
		err = DeleteConfig(filename)
		require.NoError(t, err)
	})
	t.Run("it should be failed to initialize config", func (t *testing.T) {
		conf := &errorConfig{&addr}
		_, err := InitializeConfig(conf)
		require.Error(t, err)
		require.Equal(t, addr, *conf.BackendAddr)
		err = WriteConfigToFile(conf, filename)
		require.NoError(t, err)
		err = ReadConfigFromFile(conf, filename)
		require.NoError(t, err)
		err = DeleteConfig(filename)
		require.NoError(t, err)
	})
	t.Run("it should create and delete folder", func (t *testing.T) {
		err := createFolder("banger")
		require.NoError(t, err)
		err = os.RemoveAll("banger")
		require.NoError(t, err)
	})
	t.Run("it should error when writing config", func (t *testing.T) {
		t.Setenv("HOME", "/home/somerandomhomethatsnotsupposedtobepresent")
		conf := &errorConfig{&addr}
		_, err := InitializeConfig(conf)
		require.Error(t, err)
		require.Equal(t, addr, *conf.BackendAddr)
		err = WriteConfigToFile(conf, filename)
		require.Error(t, err)
		err = ReadConfigFromFile(conf, filename)
		require.Error(t, err)
		err = DeleteConfig(filename)
		require.Error(t, err)
	})
	t.Run("it should error when $HOME env is unset", func (t *testing.T) {
		os.Unsetenv("HOME")
		conf := &errorConfig{&addr}
		_, err := InitializeConfig(conf)
		require.Error(t, err)
		require.Equal(t, addr, *conf.BackendAddr)
		err = WriteConfigToFile(conf, filename)
		require.Error(t, err)
		err = ReadConfigFromFile(conf, filename)
		require.Error(t, err)
		err = DeleteConfig(filename)
		require.Error(t, err)
		t.Setenv("HOME", tmpDir)
	})
	t.Run("it should error when marshaling a func", func (t *testing.T) {
		conf := &funcFieldConfig{func() {}}
		err := WriteConfigToFile(conf, "")
		require.Error(t, err)
	})
}
