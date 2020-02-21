package cmd

import "github.com/spf13/cobra"

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
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)
	pasteCmd.AddCommand(createCmd)
}
