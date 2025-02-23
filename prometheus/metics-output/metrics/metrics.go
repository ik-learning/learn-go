package metrics

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"sort"
	"text/template"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
	Registerer prometheus.Registerer
	Metrics    []*Metric
	mName      map[string]bool
}

type Metric struct {
	Type      string
	Namespace string
	Subsystem string
	Name      string
	Help      string
}

var (
	RegisterMetric = NewPrometheusMetrics()
)

func NewPrometheusMetrics() *PrometheusMetrics {
	reg := prometheus.WrapRegistererWith(
		prometheus.Labels{
			"version":    "test",
			"go_arch":    runtime.GOARCH,
			"go_version": runtime.Version(),
		},
		prometheus.DefaultRegisterer)
	return &PrometheusMetrics{
		Registerer: reg,
		Metrics:    []*Metric{},
		mName:      make(map[string]bool),
	}
}

var (
	erifiedAAAARecords = NewGaugeOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "verified_aaaa_records",
			Help:      "Number of DNS AAAA-records that exists both in source and registry.",
		},
	)
)

func init() {
	RegisterMetric.MustRegister(erifiedAAAARecords)
}

type GaugeMetric struct {
	Metric
	Gauge    prometheus.Gauge
	FullName string
}

func NewGaugeWithOpts(opts prometheus.GaugeOpts) GaugeMetric {
	return GaugeMetric{
		FullName: fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name),
		Metric: Metric{
			Type:      "gauge",
			Name:      opts.Name,
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		Gauge: prometheus.NewGauge(opts),
	}
}

// MustRegister TODO: only single metrics should be registered
func (m *PrometheusMetrics) MustRegister(cs interface{}) {

	switch v := cs.(type) {
	case GaugeMetric:
		fmt.Println("BINGO!!!!", v.FullName)
		return
	case prometheus.CounterOpts:
		name := fmt.Sprintf("%s_%s", v.Subsystem, v.Name)
		if _, exists := m.mName[name]; exists {
			return
		} else {
			m.mName[name] = true
		}
		m.Metrics = append(m.Metrics, &Metric{
			Type:      "counter",
			Name:      v.Name,
			Namespace: v.Namespace,
			Subsystem: v.Subsystem,
			Help:      v.Help,
		})
		m.Registerer.MustRegister(prometheus.NewCounter(v))
	case prometheus.GaugeOpts:
		name := fmt.Sprintf("%s_%s", v.Subsystem, v.Name)
		if _, exists := m.mName[name]; exists {
			return
		} else {
			m.mName[name] = true
		}
		m.Metrics = append(m.Metrics, &Metric{
			Type:      "gauge",
			Name:      v.Name,
			Namespace: v.Namespace,
			Subsystem: v.Subsystem,
			Help:      v.Help,
		})
		m.Registerer.MustRegister(prometheus.NewGauge(v))
	default:
		return // Unsupported metric type
	}
}

const markdownTemplate = `# Monitoring
<!-- THIS FILE MUST NOT BE EDITED BY HAND -->
<!-- ON NEW METRIC ADDED PLEASE RUN 'make generate-metrics-documentation' -->
<!-- markdownlint-disable MD013 -->

All metrics available for scraping are exposed on the {{backtick 1}}/metrics{{backtick 1}} endpoint. The metrics are in the Prometheus exposition format.

{{backtick 3}}sh
curl https://localhost:7979/metrics
{{backtick 3}}

## Supported Metrics

> Full metric name is constructed as follows:
> {{backtick 1}}external_dns_<subsystem>_<name>{{backtick 1}} or {{backtick 1}}external_dns.<subsystem>.<name>{{backtick 1}}.

| Name                             | Metric Type | Subsystem   |  Help                                                 |
|:---------------------------------|:------------|:------------|:------------------------------------------------------|
{{- range . }}
| {{ .Name }} | {{ .Type | capitalize }} | {{ .Subsystem }} | {{ .Help }} |
{{- end -}}
`

func (m *PrometheusMetrics) GenerateMarkdownTable() (string, error) {
	tmpl := template.Must(template.New("metrics.md.tpl").Funcs(funcMap()).Parse(markdownTemplate))

	sort.Slice(m.Metrics, func(i, j int) bool {
		if m.Metrics[i].Subsystem == m.Metrics[j].Subsystem {
			return m.Metrics[i].Name < m.Metrics[j].Name
		}
		return m.Metrics[i].Subsystem < m.Metrics[j].Subsystem
	})

	var b bytes.Buffer
	err := tmpl.Execute(&b, m.Metrics)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (m *PrometheusMetrics) WriteToFile() error {
	testPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/metrics.md", testPath)
	fmt.Printf("generate file '%s' with supported metrics\n", path)

	content, err := m.GenerateMarkdownTable()
	if err != nil {
		_ = fmt.Errorf("failed to generate markdown file '%s': %v\n", path, err.Error())
	}
	content = content + "\n"
	writeErr := writeToFile(path, content)
	return writeErr
}
