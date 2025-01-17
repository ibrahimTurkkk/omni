package log

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Promauto metrics are global.
var (
	logTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "log",
		Name:      "total",
		Help:      "Total number of log messages per level.",
	}, []string{"level"})
)
