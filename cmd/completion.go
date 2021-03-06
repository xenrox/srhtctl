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

func completeYamlFiles(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{"yaml", "yml"}, cobra.ShellCompDirectiveFilterFileExt
}

func completeNoFiles(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}
