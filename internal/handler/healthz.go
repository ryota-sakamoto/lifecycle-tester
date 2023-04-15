package handler

import (
	"net/http"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func Healthz(sm *state.StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sm.GetState().IsFailedHealthz {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}
