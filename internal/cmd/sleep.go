package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func NewSleepCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep [seconds]",
		Short: "sleep",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}

			if _, err := strconv.Atoi(args[0]); err != nil {
				return err
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			sec, _ := strconv.Atoi(args[0])
			time.Sleep(time.Second * time.Duration(sec))
		},
	}

	return cmd
}
