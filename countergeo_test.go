package metrics_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/metrics"
)

// TestCounterGeo tests the functionality of the Geo counter
func TestCounterGeo(t *testing.T) {
	geoip := &geoIpLookupProvider{ips: map[string]string{
		"127.0.0.1": "LO",
	}}
	collector := metrics.New(geoip)
	counter, err := collector.CreateCounterGeo("test", "seconds", "Hello world!")
	assert.Nil(t, err, "creating counter returned an error")

	m := collector.ListMetrics()
	assert.Equal(t, 1, len(m))
	assert.Equal(t, "test", m[0].Name)
	assert.Equal(t, "seconds", m[0].Unit)
	assert.Equal(t, "Hello world!", m[0].Help)
	assert.Equal(t, metrics.MetricTypeCounter, m[0].Type)
	assert.Equal(t, 0, len(collector.GetMetric("test")))

	counter.Increment(net.ParseIP("127.0.0.1"))
	metric := collector.GetMetric("test")
	assert.Equal(t, 1, len(metric))
	assert.Equal(t, "test", metric[0].Name)
	assert.Equal(t, float64(1), metric[0].Value)
	assert.Equal(t, map[string]string{"country": "LO"}, metric[0].Labels)

	counter.Increment(net.ParseIP("127.0.0.2"))
	metric = collector.GetMetric("test")
	var collectedMetrics []string
	for _, m := range metric {
		collectedMetrics = append(collectedMetrics, m.CombinedName())
	}
	assert.Contains(t, collectedMetrics, "test{country=\"LO\"}")
	assert.Contains(t, collectedMetrics, "test{country=\"XX\"}")
	assert.Equal(t, 2, len(metric))
	assert.Equal(t, float64(1), metric[0].Value)
	assert.Equal(t, float64(1), metric[1].Value)
}
