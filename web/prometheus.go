package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

// TotalRequests ...
var TotalRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

// ResponseStatus ...
var ResponseStatus = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

// HTTPDuration ...
var HTTPDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

// PrometheusMiddleware ...
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(HTTPDuration.WithLabelValues(path))
		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.status

		ResponseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		TotalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	})
}
