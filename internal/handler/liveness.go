package handler

import (
	"net/http"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func Liveness(sm *state.StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sm.GetState().IsFailedLiveness {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}
