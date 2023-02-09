package systemctl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	info, err := Status("cron")
	require.NoError(t, err)
	require.NotEmpty(t, info.Active)
	require.NotEmpty(t, info.MainPid)
}
