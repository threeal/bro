package systemctl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	t.Run("ValidService", func(t *testing.T) {
		info, err := Status("cron")
		require.NoError(t, err)
		require.NotEmpty(t, info.Active)
		require.NotEmpty(t, info.MainPid)
	})

	t.Run("InvalidService", func(t *testing.T) {
		_, err := Status("some-invalid-service")
		require.Error(t, err)
	})
}
