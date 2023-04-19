package cmd

import (
	"context"
	"os/signal"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func NewSleepCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep [seconds]",
		Short: "Sleep for a specified duration",
		Long: `The 'sleep' subcommand allows you to make the application sleep for a specified
duration in seconds or indefinitely using the 'infinity' keyword.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}

			if args[0] != "infinity" {
				if _, err := strconv.Atoi(args[0]); err != nil {
					return err
				}
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "infinity" {
				ctx, cancel := signal.NotifyContext(context.Background())
				defer cancel()
				<-ctx.Done()
			} else {
				sec, _ := strconv.Atoi(args[0])
				time.Sleep(time.Second * time.Duration(sec))
			}
		},
	}

	return cmd
}
