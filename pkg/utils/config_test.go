package utils

import (
	"errors"
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

func (c *testConfig) Init() error {
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

func (c *errorConfig) Init() error {
	return errors.New("init error")
}

type noFieldConfig struct {
	BackendAddr func() `json:"backend_addr"`
}

func (c *noFieldConfig) Read() error {
	return errors.New("read error")
}

func (c *noFieldConfig) Write() error {
	return errors.New("write error")
}

func (c *noFieldConfig) Init() error {
	return errors.New("init error")
}

func TestConfigIO(t *testing.T) {
	filename := "ozymandias"
	addr := ":21"
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
}

func TestConfigIOErr(t *testing.T) {
	filename := "ozymandias"
	addr := ":21"
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
}

func TestCreateFolder(t *testing.T) {
	err := createFolder("banger")
	require.NoError(t, err)
	err = os.RemoveAll("banger")
	require.NoError(t, err)
}

func TestConfigIOErrHomeDir(t *testing.T) {
	os.Setenv("HOME", "/home/somerandomhomethatsnotsupposedtobepresent")
	filename := "ozymandias"
	addr := ":21"
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
}

func TestConfigIONoHomeDir(t *testing.T) {
	os.Unsetenv("HOME")
	filename := "ozymandias"
	addr := ":21"
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
	os.Setenv("HOME", "/home/somerandomhomethatsnotsupposedtobepresent")
}

func TestConfigIOMarshalError(t *testing.T) {
	conf := &noFieldConfig{func() {}}
	err := WriteConfigToFile(conf, "")
	require.Error(t, err)
}
