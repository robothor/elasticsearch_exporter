package main

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestHasExpectedRegisteredDescs(t *testing.T) {
	exporter := NewExporter("http://example.com", "http://example.com", 10, false)

	// 1 counter for "up" plus 12 for the cluster_health metics
	expectedDescs := 1 + len(counterMetrics) + len(counterVecMetrics) + len(gaugeMetrics) + len(gaugeVecMetrics) + 12

	ch := make(chan *prometheus.Desc, 100)

	exporter.Describe(ch)
	close(ch)

	count := 0

	for _ = range ch {
		count++
	}

	if count != expectedDescs {
		t.Error("Expected", expectedDescs, "descriptions, but found", count)
	}
}
