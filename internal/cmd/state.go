package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func NewStateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state [flags]",
		Short: "Update the state of a running server",
		Long: `Update the state of a running server by providing various flags.
The command can be used to simulate failed readiness and liveness probes,
and to set a delay for the server shutdown process.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			host, err := cmd.Flags().GetString("host")
			if err != nil {
				return err
			}

			port, err := cmd.Flags().GetInt64("port")
			if err != nil {
				return err
			}

			isFailedReadiness, err := cmd.Flags().GetBool("is-failed-readiness")
			if err != nil {
				return err
			}

			isFailedLiveness, err := cmd.Flags().GetBool("is-failed-liveness")
			if err != nil {
				return err
			}

			shutdownDelaySeconds, err := cmd.Flags().GetInt64("shutdown-delay-seconds")
			if err != nil {
				return err
			}

			s := &state.State{
				IsFailedReadiness:    isFailedReadiness,
				IsFailedLiveness:     isFailedLiveness,
				ShutdownDelaySeconds: shutdownDelaySeconds,
			}
			slog.Info("update state",
				slog.Any("state", s),
			)

			return update(
				host,
				port,
				s,
			)
		},
	}

	cmd.Flags().String("host", "localhost", "The target server's hostname or IP address")
	cmd.Flags().Int64("port", 8080, "The target server's port number")
	cmd.Flags().Bool("is-failed-readiness", false, "Set to true to simulate a failed readiness probe")
	cmd.Flags().Bool("is-failed-liveness", false, "Set to true to simulate a failed liveness probe")
	cmd.Flags().Int64("shutdown-delay-seconds", 0, "Set the number of seconds to delay the server shutdown process")

	return cmd
}

func update(host string, port int64, state *state.State) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://%s:%d", host, port),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
