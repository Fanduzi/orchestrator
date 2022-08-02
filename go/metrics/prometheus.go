package metrics

import (
	"github.com/rcrowley/go-metrics"
	"time"

	prometheusmetrics "github.com/deathowl/go-metrics-prometheus"
	"github.com/openark/golib/log"
	"github.com/openark/orchestrator/go/config"
	"github.com/prometheus/client_golang/prometheus"
)

func InitPrometheusMetrics() {

	if !config.Config.EnablePrometheusMetrics {
		return
	}

	prometheusClient := prometheusmetrics.NewPrometheusProvider(metrics.DefaultRegistry, config.Config.PrometheusNamespace, config.Config.PrometheusSubsystem, prometheus.DefaultRegisterer, time.Duration(config.Config.PrometheusFlushIntervalSeconds)*time.Second)
	log.Debugf("Exporting prometheus metrics")
	go prometheusClient.UpdatePrometheusMetrics()
}
