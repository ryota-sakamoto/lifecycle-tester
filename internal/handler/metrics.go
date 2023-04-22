package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

var (
	requestsReadinessTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "requests_readiness_total",
		Help: "Total number of readiness probe requests",
	})

	requestsLivenessTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "requests_liveness_total",
		Help: "Total number of liveness probe requests",
	})
)

func init() {
	prometheus.MustRegister(
		requestsReadinessTotal,
		requestsLivenessTotal,
	)
}

func Metrics(mux *chi.Mux, sm *state.StateManager) {
	mux.Method(http.MethodGet, "/metrics", promhttp.Handler())
}
