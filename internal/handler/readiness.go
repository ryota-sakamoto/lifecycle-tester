package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func Readiness(mux *chi.Mux, sm *state.StateManager) {
	mux.Get("/readiness", readiness(sm))
}

func readiness(sm *state.StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sm.GetState().IsFailedReadiness {
			readinessRequestsTotal.WithLabelValues("503").Inc()
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			readinessRequestsTotal.WithLabelValues("200").Inc()
		}
	}
}
