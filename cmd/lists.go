package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var listsCmd = &cobra.Command{
	Use:     "lists",
	Short:   "Use the srht lists API",
	Aliases: []string{"list"},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Work on patches/patchsets",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var patchListCmd = &cobra.Command{
	Use:   "list",
	Short: "List patchsets",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.PrintPatchsets(args)
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)

	listsCmd.AddCommand(patchCmd)

	patchCmd.AddCommand(patchListCmd)
	patchListCmd.Flags().StringVarP(&api.ListName, "list", "l", "", "List name")
	patchListCmd.RegisterFlagCompletionFunc("list", completeNoFiles)
	patchListCmd.Flags().StringVarP(&api.PatchStatus, "status", "s", "proposed", "Patchset status")
	patchListCmd.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) (
		[]string, cobra.ShellCompDirective) {
		return []string{"proposed", "needs_revision", "superseded", "approved", "rejected", "applied", "all"},
			cobra.ShellCompDirectiveNoFileComp
	})
	patchListCmd.Flags().StringVarP(&api.PatchPrefix, "prefix", "p", "", "Patchset prefix")
	patchListCmd.RegisterFlagCompletionFunc("prefix", completeNoFiles)
}
