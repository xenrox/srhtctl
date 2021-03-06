package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Use the srht paste API",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new paste resource",
	Long: `Create a new paste resource.
Takes files as arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := api.PasteCreate(args)
		errorhelper.ExitError(err)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a paste resource",
	Long: `Delete a paste resource.
Takes paste SHA hashes as arguments`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.PasteDelete(args)
		errorhelper.ExitError(err)
	},
}

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Delete expired paste resources",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.PasteCleanup()
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)

	pasteCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringVarP(&api.PasteName, "name", "n", "unnamed", "Paste name")
	createCmd.PersistentFlags().StringVarP(&api.PasteVisibility, "visibility", "v", "", "Paste visibility")
	createCmd.RegisterFlagCompletionFunc("visibility", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"public", "private", "unlisted"}, cobra.ShellCompDirectiveNoFileComp
	})
	createCmd.PersistentFlags().StringVarP(&api.PasteExpiration, "expiration", "e", "", "Paste expiration in days")
	createCmd.RegisterFlagCompletionFunc("expiration", completeNoFiles)

	pasteCmd.AddCommand(deleteCmd)

	pasteCmd.AddCommand(cleanupCmd)
}
