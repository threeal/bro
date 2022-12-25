package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("FOO", "foo")
	res := GetEnvOrDefault("FOO", "default")
	require.Equal(t, "foo", res)
	os.Unsetenv("FOO")
	res = GetEnvOrDefault("FOO", "default")
	require.Equal(t, "default", res)
}
