package commands

import (
	"runtime"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "bro-backend"}
	rootCmd.AddCommand(getSpinCommand())
	if (runtime.GOOS == "linux") {
		rootCmd.AddCommand(getStatusCommand())
	}
	rootCmd.Execute()
}
