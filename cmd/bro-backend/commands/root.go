package commands

import "github.com/spf13/cobra"

func Execute() {
	var rootCmd = &cobra.Command{Use: "bro-backend"}
	rootCmd.AddCommand(getSpinCommand())
	rootCmd.Execute()
}
