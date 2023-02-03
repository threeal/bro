package commands

import "github.com/spf13/cobra"

func Execute() {
	var rootCmd = &cobra.Command{Use: "bro"}
	rootCmd.AddCommand(getEchoCommand())
	rootCmd.Execute()
}
