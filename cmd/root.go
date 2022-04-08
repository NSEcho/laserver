package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "laserver",
	Short: "Simple tracking server for lateralus",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
