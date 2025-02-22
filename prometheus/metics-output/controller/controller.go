package controller

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ik-learning/learn-go/prometheus/metics-output/metrics"
)

var (
	registryErrorsTotal = prometheus.CounterOpts{
		Namespace: "external_dns",
		Subsystem: "registry",
		Name:      "errors_total",
		Help:      "Number of Registry errors.",
	}

	sourceErrorsTotal = prometheus.CounterOpts{
		Namespace: "external_dns",
		Subsystem: "source",
		Name:      "errors_total",
		Help:      "Number of Source errors.",
	}

	sourceEndpointsTotal = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "source",
		Name:      "endpoints_total",
		Help:      "Number of Endpoints in all sources",
	}

	registryEndpointsTotal = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "registry",
		Name:      "endpoints_total",
		Help:      "Number of Endpoints in the registry",
	}
	lastSyncTimestamp = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "controller",
		Name:      "last_sync_timestamp_seconds",
		Help:      "Timestamp of last successful sync with the DNS provider",
	}

	lastReconcileTimestamp = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "controller",
		Name:      "last_reconcile_timestamp_seconds",
		Help:      "Timestamp of last attempted sync with the DNS provider",
	}

	controllerNoChangesTotal = prometheus.CounterOpts{
		Namespace: "external_dns",
		Subsystem: "controller",
		Name:      "no_op_runs_total",
		Help:      "Number of reconcile loops ending up with no changes on the DNS provider side.",
	}
	deprecatedRegistryErrors = prometheus.CounterOpts{
		Subsystem: "registry",
		Name:      "errors_total",
		Help:      "Number of Registry errors.",
	}

	deprecatedSourceErrors = prometheus.CounterOpts{
		Subsystem: "source",
		Name:      "errors_total",
		Help:      "Number of Source errors.",
	}
	registryARecords = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "registry",
		Name:      "a_records",
		Help:      "Number of Registry A records.",
	}

	registryAAAARecords = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "registry",
		Name:      "aaaa_records",
		Help:      "Number of Registry AAAA records.",
	}

	sourceARecords = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "source",
		Name:      "a_records",
		Help:      "Number of Source A records.",
	}
	sourceAAAARecords = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "source",
		Name:      "aaaa_records",
		Help:      "Number of Source AAAA records.",
	}

	verifiedARecords = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "controller",
		Name:      "verified_a_records",
		Help:      "Number of DNS A-records that exists both in source and registry.",
	}

	verifiedAAAARecords = prometheus.GaugeOpts{
		Namespace: "external_dns",
		Subsystem: "controller",
		Name:      "verified_aaaa_records",
		Help:      "Number of DNS AAAA-records that exists both in source and registry.",
	}
)

func init() {
	metrics.RegisterMetric.MustRegister(registryErrorsTotal)
	metrics.RegisterMetric.MustRegister(sourceErrorsTotal)
	metrics.RegisterMetric.MustRegister(sourceEndpointsTotal)
	metrics.RegisterMetric.MustRegister(registryEndpointsTotal)
	metrics.RegisterMetric.MustRegister(lastSyncTimestamp)
	metrics.RegisterMetric.MustRegister(lastReconcileTimestamp)
	metrics.RegisterMetric.MustRegister(deprecatedRegistryErrors)
	metrics.RegisterMetric.MustRegister(deprecatedSourceErrors)
	metrics.RegisterMetric.MustRegister(controllerNoChangesTotal)
	metrics.RegisterMetric.MustRegister(registryARecords)
	metrics.RegisterMetric.MustRegister(registryAAAARecords)
	metrics.RegisterMetric.MustRegister(sourceARecords)
	metrics.RegisterMetric.MustRegister(sourceAAAARecords)
	metrics.RegisterMetric.MustRegister(verifiedARecords)
	metrics.RegisterMetric.MustRegister(verifiedAAAARecords)
}
