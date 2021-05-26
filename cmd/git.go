package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/config"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Use the srht git API",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.PersistentFlags().StringVarP(&config.UserName, "user", "u", "", "Git user (without ~)")
}
