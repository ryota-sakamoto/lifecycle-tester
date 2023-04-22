package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests with specific endpoint and status code",
		},
		[]string{"endpoint", "code"},
	)

	readinessRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "readiness_requests_total",
		Help: "Total number of readiness probe requests",
	}, []string{"code"})

	livenessRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "liveness_requests_total",
		Help: "Total number of liveness probe requests",
	}, []string{"code"})
)

func init() {
	prometheus.MustRegister(
		httpRequestsTotal,
		readinessRequestsTotal,
		livenessRequestsTotal,
	)

	readinessRequestsTotal.WithLabelValues("200")
	readinessRequestsTotal.WithLabelValues("503")

	livenessRequestsTotal.WithLabelValues("200")
	livenessRequestsTotal.WithLabelValues("503")
}

func Metrics(mux *chi.Mux, sm *state.StateManager) {
	mux.Method(http.MethodGet, "/metrics", promhttp.Handler())
}

func IncHttpRequestsTotal(endpoint string, code int) {
	httpRequestsTotal.WithLabelValues(endpoint, strconv.Itoa(code)).Inc()
}
