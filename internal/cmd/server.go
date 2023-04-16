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
		Short: "Run http server",
		Run:   runServer,
	}

	return cmd
}

func runServer(cmd *cobra.Command, args []string) {
	c, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))
	slog.Info("start http server",
		slog.Any("config", c),
	)

	sm := state.NewStateManager()
	sm.SetState(state.State{
		IsFailedReadiness:    false,
		IsFailedLiveness:     false,
		ShutdownDelaySeconds: c.ShutdownDelaySeconds,
	})

	mux := chi.NewRouter()
	mux.Use(middleware.Logging(c.EnableHealthLog))

	mux.Get("/", handler.GetIndex(sm))
	mux.Post("/", handler.PostIndex(sm))
	mux.Get("/readiness", handler.Readiness(sm))
	mux.Get("/liveness", handler.Liveness(sm))

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