package cmd

import (
	"github.com/spf13/cobra"
)

func NewStateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: "state",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := cmd.Flags().GetInt64("shutdown-delay-seconds")
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().Int64("shutdown-delay-seconds", 0, "")

	return cmd
}
