package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/metrics"
	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func Liveness(mux *chi.Mux, sm *state.StateManager) {
	mux.Get("/liveness", liveness(sm))
}

func liveness(sm *state.StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sm.GetState().IsFailedLiveness {
			metrics.LivenessRequestsTotal.WithLabelValues("503").Inc()
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			metrics.LivenessRequestsTotal.WithLabelValues("200").Inc()
		}
	}
}
