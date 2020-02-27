package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
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
	Short: "Deploy build manifest(s)",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildDeploy(args)
		errorhelper.ExitError(err)
	},
}

var resubmitCmd = &cobra.Command{
	Use:   "resubmit",
	Short: "Resubmit a build",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildResubmit(args)
		errorhelper.ExitError(err)
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information about a job by its ID",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildInformation(args)
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.AddCommand(deployCmd)

	buildCmd.AddCommand(resubmitCmd)
	resubmitCmd.PersistentFlags().BoolVarP(&api.BuildEdit, "edit", "e", false, "Edit manifest")

	buildCmd.AddCommand(infoCmd)
}
