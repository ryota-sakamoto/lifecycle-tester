package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/state"
)

var (
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
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

	httpConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_connections",
			Help: "Number of active HTTP connections",
		},
	)
)

func init() {
	prometheus.MustRegister(
		httpRequestDuration,
		readinessRequestsTotal,
		livenessRequestsTotal,
		httpConnections,
	)

	readinessRequestsTotal.WithLabelValues("200")
	readinessRequestsTotal.WithLabelValues("503")

	livenessRequestsTotal.WithLabelValues("200")
	livenessRequestsTotal.WithLabelValues("503")
}

func Metrics(mux *chi.Mux, sm *state.StateManager) {
	mux.Method(http.MethodGet, "/metrics", promhttp.Handler())
}

func RecordHttpRequestDuration(endpoint string, code int, t time.Duration) {
	httpRequestDuration.WithLabelValues(endpoint, strconv.Itoa(code)).Observe(t.Seconds())
}

func TrackHttpConnections() func() {
	httpConnections.Inc()
	return func() {
		httpConnections.Dec()
	}
}
