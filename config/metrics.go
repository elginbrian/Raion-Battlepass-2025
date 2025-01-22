package config

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	httpRequestSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_size_bytes",
			Help:    "Size of HTTP requests in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 5),
		},
		[]string{"method", "endpoint"},
	)

	httpResponseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 5),
		},
		[]string{"method", "endpoint"},
	)

	httpConcurrentRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_concurrent_requests",
			Help: "Number of concurrent HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	once sync.Once
)

func InitMetrics() {
	once.Do(func() {
		prometheus.MustRegister(httpRequestsTotal)
		prometheus.MustRegister(httpRequestDuration)
		prometheus.MustRegister(httpRequestSize)
		prometheus.MustRegister(httpResponseSize)
		prometheus.MustRegister(httpConcurrentRequests)
	})
}

func PrometheusMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	httpConcurrentRequests.WithLabelValues(c.Method(), normalizePath(c.Path())).Inc()

	recorder := &statusRecorder{statusCode: http.StatusOK}

	err := c.Next()

	duration := time.Since(start).Seconds()
	normalizedEndpoint := normalizePath(c.Path())

	httpRequestsTotal.WithLabelValues(c.Method(), normalizedEndpoint, http.StatusText(recorder.statusCode)).Inc()
	httpRequestDuration.WithLabelValues(c.Method(), normalizedEndpoint).Observe(duration)
	httpRequestSize.WithLabelValues(c.Method(), normalizedEndpoint).Observe(float64(c.Request().Header.ContentLength()))
	httpResponseSize.WithLabelValues(c.Method(), normalizedEndpoint).Observe(float64(recorder.size))

	httpConcurrentRequests.WithLabelValues(c.Method(), normalizedEndpoint).Dec()

	return err
}

func normalizePath(path string) string {
	segments := strings.Split(path, "/")
	for i, segment := range segments {
		if isID(segment) {
			segments[i] = "{id}"
		}
	}
	return strings.Join(segments, "/")
}

func isID(segment string) bool {
	return strings.Contains(segment, "-") || strings.ContainsAny(segment, "0123456789")
}

type statusRecorder struct {
	*fiber.Ctx
	statusCode int
	size       int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.Ctx.Response().SetStatusCode(code)
}

func (rec *statusRecorder) Write(b []byte) (int, error) {
	size, err := rec.Ctx.Response().BodyWriter().Write(b)
	rec.size += size
	return size, err
}