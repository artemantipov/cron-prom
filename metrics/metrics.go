package metrics

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metricsPort  = getEnv("METRIC_PORT", "1221")
	metricsURL   = getEnv("METRICS_URL", "/metrics")
	metricPrefix = getEnv("METRIC_PREFIX", "")
	// JobsFailed counter
	JobsFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%vcron_jobs_failed", metricPrefix),
		Help: "Number of failed jobs",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(JobsFailed)
}

// Start prometheus metrics endpoint
func Start() {
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle(metricsURL, promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", metricsPort), nil))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
