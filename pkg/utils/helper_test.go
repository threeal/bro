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
	t.Run("GetJSONFields should return slices of string with length of 3", func(t *testing.T) {
		res := GetJSONFields(helperConfig{"yes", "yes", "yes", "yes"})
		require.Len(t, res, 3)
	})

	t.Run("GetStructValueByJSON should return value of yes", func(t *testing.T) {
		res := GetStructValueByJSON(helperConfig{"yes", "yes", "yes", "yes"}, "backend_addr")
		require.Equal(t, "yes", res)
	})

	t.Run("GetStructValueByJSON should return empty string", func(t *testing.T) {
		res := GetStructValueByJSON(helperConfig{"yes", "yes", "yes", "yes"}, "not_exist")
		require.Empty(t, res)
	})

	t.Run("SetStructValueByJSON should set the field correctly", func(t *testing.T) {
		c := &helperConfig{"yes", "yes", "yes", "yes"}
		SetStructValueByJSON(c, "backend_addr", "banger")
		require.Equal(t, c.BackendAddr, "banger")
	})
}
