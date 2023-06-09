package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

type Response struct {
	Hostname string      `json:"hostname"`
	Request  HTTPRequest `json:"request"`
	State    state.State `json:"state"`
}

type HTTPRequest struct {
	Header     http.Header `json:"header"`
	Host       string      `json:"host"`
	Method     string      `json:"method"`
	RequestURI string      `json:"request_uri"`
	RemoteAddr string      `json:"remote_addr"`
}

func Index(mux *chi.Mux, sm *state.StateManager) {
	mux.Handle("/*", getIndex(sm))
	mux.Post("/", postIndex(sm))
}

func getIndex(sm *state.StateManager) http.HandlerFunc {
	hostname, _ := os.Hostname()

	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, Response{
			Hostname: hostname,
			Request:  PickRequest(r),
			State:    sm.GetState(),
		})
	}
}

func PickRequest(r *http.Request) HTTPRequest {
	return HTTPRequest{
		Header:     r.Header,
		Host:       r.Host,
		Method:     r.Method,
		RequestURI: r.RequestURI,
		RemoteAddr: r.RemoteAddr,
	}
}

func postIndex(sm *state.StateManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var stateReq state.State
		if err := json.NewDecoder(r.Body).Decode(&stateReq); err != nil {
			slog.Warn("failed to parse state request",
				slog.Any("err", err),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		slog.Info("state change",
			slog.Any("state", stateReq),
		)
		sm.SetState(stateReq)
	}
}
