package cmd

import (
	"git.xenrox.net/~xenrox/srhtctl/api"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Use the srht todo API",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var ticketCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Work on tickets",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var ticketListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tickets",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.PrintTickets(args)
		errorhelper.ExitError(err)
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	todoCmd.AddCommand(ticketCmd)

	ticketCmd.AddCommand(ticketListCmd)
	ticketListCmd.Flags().StringVarP(&api.TicketStatus, "status", "s", "reported", "Ticket status")
	ticketListCmd.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) (
		[]string, cobra.ShellCompDirective) {
		return []string{"reported", "confirmed", "in_progress", "pending", "resolved", "all"},
			cobra.ShellCompDirectiveNoFileComp
	})
	ticketListCmd.Flags().StringVarP(&api.TrackerName, "tracker", "t", "", "Tracker name")
	ticketListCmd.RegisterFlagCompletionFunc("tracker", completeNoFiles)
}
