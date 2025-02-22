package cache

import (
	"github.com/ik-learning/learn-go/prometheus/metics-output/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	cachedRecordsCallsTotal = prometheus.CounterOpts{
		Namespace: "external_dns",
		Subsystem: "provider",
		Name:      "cache_records_calls",
		Help:      "Number of calls to the provider cache Records list.",
	}

	cachedApplyChangesCallsTotal = prometheus.CounterOpts{
		Namespace: "external_dns",
		Subsystem: "provider",
		Name:      "cache_apply_changes_calls",
		Help:      "Number of calls to the provider cache ApplyChanges.",
	}
)

func init() {
	metrics.RegisterMetric.MustRegister(cachedRecordsCallsTotal)
	metrics.RegisterMetric.MustRegister(cachedApplyChangesCallsTotal)
}
