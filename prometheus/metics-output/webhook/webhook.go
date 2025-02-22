package webhook

import (
	"github.com/ik-learning/learn-go/prometheus/metics-output/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	recordsErrorsGauge = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "webhook_provider",
		Name:      "records_errors_total",
		Help:      "Errors with Records method",
	}

	recordsRequestsGauge = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "webhook_provider",
		Name:      "records_requests_total",
		Help:      "Requests with Records method",
	}
	applyChangesErrorsGauge = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "webhook_provider",
		Name:      "applychanges_errors_total",
		Help:      "Errors with ApplyChanges method",
	}

	applyChangesRequestsGauge = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "webhook_provider",
		Name:      "applychanges_requests_total",
		Help:      "Requests with ApplyChanges method",
	}

	adjustEndpointsErrorsGauge = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "webhook_provider",
		Name:      "adjustendpoints_errors_total",
		Help:      "Errors with AdjustEndpoints method",
	}
	adjustEndpointsRequestsGauge = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "webhook_provider",
		Name:      "adjustendpoints_requests_total",
		Help:      "Requests with AdjustEndpoints method",
	}
)

func init() {
	metrics.RegisterMetric.MustRegister(recordsErrorsGauge)
	metrics.RegisterMetric.MustRegister(recordsRequestsGauge)
	metrics.RegisterMetric.MustRegister(applyChangesErrorsGauge)
	metrics.RegisterMetric.MustRegister(applyChangesRequestsGauge)
	metrics.RegisterMetric.MustRegister(adjustEndpointsErrorsGauge)
	metrics.RegisterMetric.MustRegister(adjustEndpointsRequestsGauge)
}
