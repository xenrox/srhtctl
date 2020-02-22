package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
	"github.com/spf13/cobra"
)

var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Use the srht paste API",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new paste resource",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.PasteCreate()
		helpers.PrintError(err)
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)

	pasteCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringVarP(&api.PasteFile, "file", "f", "", "File to upload")
	createCmd.PersistentFlags().StringVarP(&api.PasteName, "name", "n", "unnamed", "Name for paste file")
	createCmd.PersistentFlags().StringVarP(&api.PasteVisibility, "visibility", "v", "", "Paste visibility")
	createCmd.PersistentFlags().StringVarP(&api.PasteExpiration, "expiration", "e", "", "Paste expiration in days")
}
