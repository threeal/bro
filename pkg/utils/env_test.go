package utils

import "testing"

func TestGetEnvOrDefault(t *testing.T) {
	res := GetEnvOrDefault("PATH", "foo")
	if res == "foo" {
		t.Errorf("not expected 'foo' result")
	}
	res = GetEnvOrDefault("FOO", "foo")
	if res != "foo" {
		t.Errorf("expected 'foo', got: %s", res)
	}
}
