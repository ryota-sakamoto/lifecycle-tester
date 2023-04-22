package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func Liveness(mux *chi.Mux, sm *state.StateManager) {
	mux.Get("/liveness", liveness(sm))
}

func liveness(sm *state.StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sm.GetState().IsFailedLiveness {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}
