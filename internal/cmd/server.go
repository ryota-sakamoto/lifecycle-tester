package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/config"
	"github.com/ryota-sakamoto/lifecycle-tester/internal/handler"
	"github.com/ryota-sakamoto/lifecycle-tester/internal/middleware"
	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run the HTTP server",
		Long: `The 'server' subcommand starts an HTTP server to provide an interface for
updating the application's state and handling container lifecycle events.`,
		Run: runServer,
	}

	return cmd
}

func runServer(cmd *cobra.Command, args []string) {
	c, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	slog.Info("start http server",
		slog.Any("config", c),
	)

	sm := state.NewStateManager()
	sm.SetState(state.State{
		IsFailedReadiness:    c.ReadinessDelaySeconds > 0,
		IsFailedLiveness:     c.LivenessDelaySeconds > 0,
		ShutdownDelaySeconds: c.ShutdownDelaySeconds,
	})
	if c.ReadinessDelaySeconds > 0 {
		go func() {
			time.Sleep(time.Second * time.Duration(c.ReadinessDelaySeconds))
			current := sm.GetState()
			current.IsFailedReadiness = false
			sm.SetState(current)
		}()
	}
	if c.LivenessDelaySeconds > 0 {
		go func() {
			time.Sleep(time.Second * time.Duration(c.LivenessDelaySeconds))
			current := sm.GetState()
			current.IsFailedLiveness = false
			sm.SetState(current)
		}()
	}

	mux := chi.NewRouter()
	mux.Use(middleware.Logging(c.DisableHealthLog))

	handler.Index(mux, sm)
	handler.Readiness(mux, sm)
	handler.Liveness(mux, sm)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		slog.Info("starting http server")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()
	<-ctx.Done()

	slog.Info("stopping http server")
	if sm.GetState().ShutdownDelaySeconds > 0 {
		slog.Info("delaying shutdown",
			slog.Int64("shutdown_delay_seconds", sm.GetState().ShutdownDelaySeconds),
		)
		time.Sleep(time.Second * time.Duration(sm.GetState().ShutdownDelaySeconds))
	}

	if err := server.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	slog.Info("stopped http server")
}
