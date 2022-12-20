package utils

import (
	"os"
	"testing"
)

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("FOO", "foo")
	res := GetEnvOrDefault("FOO", "default")
	if res != "foo" {
		t.Errorf("expected 'foo', got: %s", res)
	}
	os.Unsetenv("FOO")
	res = GetEnvOrDefault("FOO", "default")
	if res != "default" {
		t.Errorf("expected 'default', got: %s", res)
	}
}
