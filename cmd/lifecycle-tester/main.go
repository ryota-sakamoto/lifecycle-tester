package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/handler"
	"github.com/ryota-sakamoto/lifecycle-tester/internal/middleware"
)

func main() {
	sm := handler.NewStateManager()

	mux := chi.NewRouter()
	mux.Use(middleware.Logging)

	mux.Get("/", handler.GetIndex(sm))
	mux.Post("/", handler.PostIndex(sm))
	mux.Get("/healthz", handler.Healthz(sm))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	go func() {
		slog.Info("start http server")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
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
