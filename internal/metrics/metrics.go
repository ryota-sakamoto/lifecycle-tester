package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint", "code"},
	)

	ReadinessRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "readiness_requests_total",
		Help: "Total number of readiness probe requests",
	}, []string{"code"})

	LivenessRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "liveness_requests_total",
		Help: "Total number of liveness probe requests",
	}, []string{"code"})

	HttpConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_connections",
			Help: "Number of active HTTP connections",
		},
	)
)

func init() {
	prometheus.MustRegister(
		HttpRequestDuration,
		ReadinessRequestsTotal,
		LivenessRequestsTotal,
		HttpConnections,
	)

	ReadinessRequestsTotal.WithLabelValues("200")
	ReadinessRequestsTotal.WithLabelValues("503")

	LivenessRequestsTotal.WithLabelValues("200")
	LivenessRequestsTotal.WithLabelValues("503")
}
