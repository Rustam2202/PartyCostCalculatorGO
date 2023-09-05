package metrics

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func RecordMetrics() gin.HandlerFunc {
	requestCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "route"},
	)
	responseTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_response_time_seconds",
		Help:    "HTTP response time in seconds",
		Buckets: []float64{0.1, 0.5, 1, 2, 5},
	}, []string{"method", "route"})

	prometheus.MustRegister(requestCount, responseTime)

	return func(ctx *gin.Context) {
		start := time.Now()
		duration := time.Since(start).Seconds()
		labels := prometheus.Labels{
			"method": ctx.Request.Method,
			"route":  ctx.Request.URL.Path,
		}
		requestCount.With(labels).Inc()
		responseTime.With(labels).Observe(duration)
	}
}
