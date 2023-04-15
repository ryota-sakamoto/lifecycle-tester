package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "lifecycle-tester",
		Run: runHelp,
	}
	cmd.AddCommand(NewSleepCommand())
	cmd.AddCommand(NewServerCommand())

	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
