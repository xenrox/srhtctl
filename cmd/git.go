package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Use the srht git API",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Create annotations",
	Long: `Create annotations.
Takes one annotations.json as argument.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.GitAnnotate(args)
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.PersistentFlags().StringVarP(&config.UserName, "user", "u", "", "Git user (without ~)")

	gitCmd.AddCommand(annotateCmd)
	annotateCmd.Flags().StringVarP(&api.GitRepoName, "repo", "r", "", "Git repository name")
	annotateCmd.MarkFlagRequired("repo")
}
