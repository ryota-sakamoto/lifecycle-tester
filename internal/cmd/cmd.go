package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lifecycle-tester [command]",
		Short: "A CLI tool for testing container lifecycle events",
		Long: `lifecycle-tester is a CLI tool designed to help test container lifecycle events, such as
readiness and liveness probes. It provides various subcommands for different functionalities.`,
		Run: runHelp,
	}
	cmd.AddCommand(NewSleepCommand())
	cmd.AddCommand(NewServerCommand())
	cmd.AddCommand(NewStateCommand())

	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
