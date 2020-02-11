package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (

	// prometheus
	inFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "in_flight_requests",
		Help: "A gauge of requests currently being served by the wrapped handler.",
	})

	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "A counter for requests to the wrapped handler.",
		},
		[]string{"code", "method"},
	)

	// duration is partitioned by the HTTP method and handler. It uses custom
	// buckets based on the expected request duration.
	duration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "A histogram of latencies for requests.",
			Buckets: []float64{.25, .5, 1, 2.5, 5, 10},
		},
		[]string{"handler", "method"},
	)

	// responseSize has no labels, making it a zero-dimensional
	// ObserverVec.
	responseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "response_size_bytes",
			Help:    "A histogram of response sizes for requests.",
			Buckets: []float64{200, 500, 900, 1500},
		},
		[]string{},
	)
)

func init() {
	prometheus.MustRegister(inFlightGauge, counter, duration, responseSize)
}

func CreateInstrumentedHandler(realHandler http.Handler, handlerName string) http.Handler {
	return handlers.LoggingHandler(os.Stdout,
		promhttp.InstrumentHandlerInFlight(inFlightGauge,
			promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": handlerName}),
				promhttp.InstrumentHandlerCounter(counter,
					promhttp.InstrumentHandlerResponseSize(responseSize, realHandler),
				),
			),
		),
	)
}
