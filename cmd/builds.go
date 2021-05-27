package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var buildsCmd = &cobra.Command{
	Use:     "builds",
	Short:   "Use the srht build API",
	Aliases: []string{"build"},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy build manifest(s)",
	Long: `Deploy build manifest(s)
Takes yml manifests as arguments.`,
	Args:              cobra.MinimumNArgs(1),
	ValidArgsFunction: completeYamlFiles,
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildDeploy(args)
		errorhelper.ExitError(err)
	},
}

var resubmitCmd = &cobra.Command{
	Use:   "resubmit",
	Short: "Resubmit a build",
	Long: `Resubmit a build.
Takes one job ID as argument.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildResubmit(args)
		errorhelper.ExitError(err)
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information about a job by its ID",
	Long: `Get information about a job by its ID.
Takes one job ID as argument.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.BuildInformation(args)
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(buildsCmd)

	buildsCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().StringVarP(&api.BuildNote, "note", "n", "", "Build note")
	deployCmd.PersistentFlags().StringSliceVarP(&api.BuildTags, "tags", "t", nil, "Comma seperated string of tags")

	buildsCmd.AddCommand(resubmitCmd)
	resubmitCmd.PersistentFlags().BoolVarP(&api.BuildEdit, "edit", "e", false, "Edit manifest")

	buildsCmd.AddCommand(infoCmd)
}
