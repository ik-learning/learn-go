package main

import (
	"fmt"

	"github.com/ik-learning/learn-go/prometheus/metics-output/metrics"
	"github.com/prometheus/client_golang/prometheus"

	_ "github.com/ik-learning/learn-go/prometheus/metics-output/cache"
	_ "github.com/ik-learning/learn-go/prometheus/metics-output/controller"
	_ "github.com/ik-learning/learn-go/prometheus/metics-output/webhook"
)

const markdownTemplate = `# Monitoring
<!-- THIS FILE MUST NOT BE EDITED BY HAND -->
<!-- ON NEW METRIC ADDED PLEASE RUN 'make generate-metrics-documentation' -->
<!-- markdownlint-disable MD013 -->

| Name | Metric Type  | Help
| :------ | :----------- |
{{- range . }}
| {{ .Name }} | {{ .Description }} |
{{- end -}}
`

func main() {
	reg := prometheus.DefaultRegisterer
	fmt.Println(reg)
	fmt.Println(len(metrics.RegisterMetric.Metrics))
}
