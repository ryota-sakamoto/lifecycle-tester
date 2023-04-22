package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func Metrics(mux *chi.Mux, sm *state.StateManager) {
	mux.Method(http.MethodGet, "/metrics", promhttp.Handler())
}
