package cmd

import (
	"fmt"
	"os"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "srhtctl",
	Short: "CLI for using the sourcehut API",
	Long: `srhtctl is a CLI for using the sourcehut API.
The project is hosted at https://git.xenrox.net/~xenrox/srhtctl`,
}

// Execute is the entrypoint for the main fucntion.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVar(&config.ConfigPath, "config", "", "Path to config.ini")
}
