package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "ssm",
	Short: "ssm is a CLI tool for file management.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
