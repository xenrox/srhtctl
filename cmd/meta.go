package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Use the srht meta API",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get profile information",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := api.MetaGetProfile()
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(metaCmd)

	metaCmd.AddCommand(profileCmd)
	profileCmd.PersistentFlags().BoolVarP(&api.MetaEdit, "edit", "e", false, "Edit profile information")
}
