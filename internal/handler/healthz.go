package handler

import (
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	if state.IsFailedHealthz {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
