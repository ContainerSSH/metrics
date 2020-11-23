package metrics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/metrics"
)

// TestCounter tests the functionality of counters
func TestCounter(t *testing.T) {
	collector := metrics.New(&geoIpLookupProvider{})
	counter, err := collector.CreateCounter("test", "seconds", "Hello world!")
	assert.Nil(t, err, "creating counter returned an error")

	m := collector.ListMetrics()
	assert.Equal(t, 1, len(m))
	assert.Equal(t, "test", m[0].Name)
	assert.Equal(t, "seconds", m[0].Unit)
	assert.Equal(t, "Hello world!", m[0].Help)
	assert.Equal(t, metrics.MetricTypeCounter, m[0].Type)
	assert.Equal(t, 0, len(collector.GetMetric("test")))

	counter.Increment()
	metric := collector.GetMetric("test")
	assert.Equal(t, 1, len(metric))
	assert.Equal(t, "test", metric[0].Name)
	assert.Equal(t, float64(1), metric[0].Value)
	assert.Equal(t, 0, len(metric[0].Labels))

	counter.Increment()
	metric = collector.GetMetric("test")
	assert.Equal(t, 1, len(metric))
	assert.Equal(t, "test", metric[0].Name)
	assert.Equal(t, float64(2), metric[0].Value)
	assert.Equal(t, 0, len(metric[0].Labels))

	err = counter.IncrementBy(2)
	assert.Nil(t, err, "incrementing a counter returned an error")
	metric = collector.GetMetric("test")
	assert.Equal(t, 1, len(metric))
	assert.Equal(t, "test", metric[0].Name)
	assert.Equal(t, float64(4), metric[0].Value)
	assert.Equal(t, 0, len(metric[0].Labels))

	err = counter.IncrementBy(-1)
	assert.EqualError(
		t,
		err,
		metrics.CounterCannotBeIncrementedByNegative.Error(),
		"incrementing a counter by negative number did not return an error",
	)
}
