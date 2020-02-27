package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionZSHCmd = &cobra.Command{
	Use:    "completionZSH",
	Short:  "Generates zsh completion scripts",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}

var completionBASHCmd = &cobra.Command{
	Use:    "completionBASH",
	Short:  "Generates bash completion scripts",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionZSHCmd)
	rootCmd.AddCommand(completionBASHCmd)
}
