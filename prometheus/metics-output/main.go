package main

import (
	"fmt"

	_ "github.com/ik-learning/learn-go/prometheus/metics-output/cache"
	_ "github.com/ik-learning/learn-go/prometheus/metics-output/controller"
	"github.com/ik-learning/learn-go/prometheus/metics-output/metrics"
	_ "github.com/ik-learning/learn-go/prometheus/metics-output/webhook"
)

func main() {
	fmt.Println(len(metrics.RegisterMetric.Metrics))

	err := metrics.RegisterMetric.WriteToFile()
	fmt.Println(err)

	// http.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
	// 	EnableOpenMetrics: true,
	// }))
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
