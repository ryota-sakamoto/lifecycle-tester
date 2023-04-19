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
		Use:   "sleep [seconds|infinity]",
		Short: "Sleep for a specified number of seconds or indefinitely",
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
