package handler_test

import (
	"net/http/httptest"

	"github.com/go-chi/chi/v5"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/handler"
	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

func setupTestServer(sm *state.StateManager) *httptest.Server {
	mux := chi.NewRouter()
	handler.Index(mux, sm)
	handler.Readiness(mux, sm)
	handler.Liveness(mux, sm)

	return httptest.NewServer(mux)
}
