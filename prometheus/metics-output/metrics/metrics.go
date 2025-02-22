package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
	Metrics []*Metric
}

type Metric struct {
	Type      string
	Namespace string
	Subsystem string
	Name      string
	Help      string
}

var (
	RegisterMetric = PrometheusMetrics{
		Metrics: []*Metric{},
	}
)

func (m *PrometheusMetrics) MustRegister(cs interface{}) {
	metricType := ""
	var opts prometheus.Opts

	switch v := cs.(type) {
	case prometheus.CounterOpts:
		metricType = "counter"
		m.Metrics = append(m.Metrics, &Metric{
			Type:      metricType,
			Name:      opts.Name,
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		})
		prometheus.DefaultRegisterer.MustRegister(prometheus.NewCounter(v))
	case prometheus.GaugeOpts:
		metricType = "gauge"
		m.Metrics = append(m.Metrics, &Metric{
			Type:      metricType,
			Name:      opts.Name,
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		})
		prometheus.DefaultRegisterer.MustRegister(prometheus.NewGauge(v))
	default:
		return // Unsupported metric type
	}
}
