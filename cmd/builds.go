package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Use the srht build API",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a build file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildDeploy(args)
		helpers.ExitError(err)
	},
}

var resubmitCmd = &cobra.Command{
	Use:   "resubmit",
	Short: "Resubmit a build",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildResubmit(args)
		helpers.ExitError(err)
	}}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.AddCommand(deployCmd)
	buildCmd.AddCommand(resubmitCmd)
}
