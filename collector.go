package metrics

import (
	"errors"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"
)

// MetricType is the enum for tye types of metrics supported
type MetricType string

const (
	// MetricTypeCounter is a data type that contains ever increasing numbers from the start of the server.
	MetricTypeCounter MetricType = "counter"

	// MetricTypeGauge is a metric type that can increase or decrease depending on the current value.
	MetricTypeGauge MetricType = "gauge"
)

// Metric is a descriptor for metrics.
type Metric struct {
	// Name is the name for the metric.
	Name string

	// Help is the help text for this metric.
	Help string

	// Unit describes the unit of the metric.
	Unit string

	// Created describes the time the metric was created. This is important for counters.
	Created time.Time

	// Type describes how the metric behaves.
	Type MetricType
}

// String formats a metric as the OpenMetrics metadata
func (metric Metric) String() string {
	return fmt.Sprintf(
		"# HELP %s %s\n"+
			"# UNIT %s %s\n"+
			"# TYPE %s %s\n",
		metric.Name,
		metric.Help,
		metric.Name,
		metric.Unit,
		metric.Name,
		metric.Type)
}

// MetricValue is a structure that contains a value for a specific metric name and set of values.
type MetricValue struct {
	// Name contains the name of the value.
	Name string

	// Labels contains a key-value map of labels to which the Value is specific.
	Labels map[string]string

	// Value contains the specific value stored.
	Value float64
}

// CombinedName returns the name and labels combined.
func (metricValue MetricValue) CombinedName() string {
	var labelList []string

	keys := make([]string, 0, len(metricValue.Labels))
	for k := range metricValue.Labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		// TODO escaping
		labelList = append(labelList, k+"=\""+metricValue.Labels[k]+"\"")
	}

	var labels string
	if len(labelList) > 0 {
		labels = "{" + strings.Join(labelList, ",") + "}"
	} else {
		labels = ""
	}

	return metricValue.Name + labels
}

// String creates a string out of the name, labels, and value.
func (metricValue MetricValue) String() string {
	return fmt.Sprintf("%s %f\n", metricValue.CombinedName(), metricValue.Value)
}

// MetricAlreadyExists is an error that is returned from the Create functions when the metric already exists.
var MetricAlreadyExists = errors.New("the specified metric already exists")

// CounterCannotBeIncrementedByNegative is an error returned by counters when they are incremented with a negative
//                                      number.
var CounterCannotBeIncrementedByNegative = errors.New("a counter cannot be incremented by a negative number")

// Collector is the main interface for interacting with the metrics collector.
type Collector interface {
	// CreateCounter creates a monotonic (increasing) counter with the specified name and help text.
	CreateCounter(name string, unit string, help string) (SimpleCounter, error)

	// CreateCounterGeo creates a monotonic (increasing) counter that is labeled with the country from the GeoIP lookup
	//                  with the specified name and help text.
	CreateCounterGeo(name string, unit string, help string) (SimpleGeoCounter, error)

	// CreateGauge creates a freely modifiable numeric gauge with the specified name and help text.
	CreateGauge(name string, unit string, help string) (SimpleGauge, error)

	// CreateGaugeGeo creates a freely modifiable numeric gauge that is labeled with the country from the GeoIP lookup
	//                with the specified name and help text.
	CreateGaugeGeo(name string, unit string, help string) (SimpleGeoGauge, error)

	// ListMetrics returns a list of metrics metadata stored in the collector.
	ListMetrics() []Metric

	// GetMetric returns a set of values with labels for a specified metric name.
	GetMetric(name string) []MetricValue

	// String returns a Prometheus/OpenMetrics-compatible document with all metrics.
	String() string
}

// SimpleCounter is a simple counter that can only be incremented.
type SimpleCounter interface {
	// Increment increments the counter by 1
	Increment()

	// IncrementBy increments the counter by the specified number. Only returns an error if the passed by parameter is
	//             negative.
	IncrementBy(by float64) error
}

// SimpleGeoCounter is a simple counter that can only be incremented and is labeled with the country from a GeoIP
//                  lookup.
type SimpleGeoCounter interface {
	// Increment increments the counter for the country from the specified ip by 1.
	Increment(ip net.IP)

	// IncrementBy increments the counter for the country from the specified ip by the specified value.
	//             Only returns an error if the passed by parameter is negative.
	IncrementBy(ip net.IP, by float64) error
}

// SimpleGauge is a metric that can be incremented and decremented.
type SimpleGauge interface {
	// Increment increments the counter by 1
	Increment()

	// IncrementBy increments the counter by the specified number.
	IncrementBy(by float64)

	// Decrement decreases the metric by 1.
	Decrement()

	// Decrement decreases the metric by the specified value.
	DecrementBy(by float64)

	// Set sets the value of the metric to an exact value.
	Set(value float64)
}

// SimpleGeoGauge is a metric that can be incremented and decremented and is labeled by the country from a GeoIP lookup.
type SimpleGeoGauge interface {
	// Increment increments the counter for the country from the specified ip by 1.
	Increment(ip net.IP)

	// IncrementBy increments the counter for the country from the specified ip by the specified value.
	IncrementBy(ip net.IP, by float64)

	// Decrement decreases the value for the country looked up from the specified IP by 1.
	Decrement(ip net.IP)

	// DecrementBy decreases the value for the country looked up from the specified IP by the specified value.
	DecrementBy(ip net.IP, by float64)

	// Set sets the value of the metric for the country looked up from the specified IP.
	Set(ip net.IP, value float64)
}
