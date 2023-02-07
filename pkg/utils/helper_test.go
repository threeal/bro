package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type helperConfig struct {
	BackendAddr  string `json:"backend_addr"`
	Foo          string
	SecondField  string `json:"-"`
	AnotherField string `json:","`
}

func TestHelper(t *testing.T) {
	require.True(t, StringInSlice("a", []string{"a", "b"}))
	require.False(t, StringInSlice("c", []string{"a", "b"}))
	res := GetJSONFields(helperConfig{"yes", "yes", "yes", "yes"})
	require.Len(t, res, 3)
}
