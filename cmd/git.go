package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Use the srht git API",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Create annotations",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.GitAnnotate(args)
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)

	gitCmd.AddCommand(annotateCmd)
	gitCmd.PersistentFlags().StringVarP(&api.GitUserName, "user", "u", "", "Git user (without ~)")
	gitCmd.PersistentFlags().StringVarP(&api.GitRepoName, "repo", "r", "", "Git repository name")
}