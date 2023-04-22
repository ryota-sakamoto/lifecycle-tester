package middleware

import (
	"net/http"
	"time"

	"golang.org/x/exp/slog"

	"github.com/ryota-sakamoto/lifecycle-tester/internal/handler"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (s *statusResponseWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

func Logging(disableHealthLog bool) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			s := &statusResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			track := handler.TrackHttpConnections()
			defer track()

			h.ServeHTTP(s, r)

			elapsed := time.Since(start)
			handler.RecordHttpRequestDuration(r.RequestURI, s.statusCode, elapsed)

			if disableHealthLog && (r.RequestURI == "/readiness" || r.RequestURI == "/liveness") {
				return
			}

			slog.Info("access log",
				slog.Int("status", s.statusCode),
				slog.String("elapsed", elapsed.String()),
				slog.Any("request", handler.PickRequest(r)),
			)
		})
	}
}