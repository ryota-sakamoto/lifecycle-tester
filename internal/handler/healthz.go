package handler

import (
	"net/http"
)

func Healthz(sm *StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sm.GetState().IsFailedHealthz {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}
