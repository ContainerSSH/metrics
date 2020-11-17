package metrics

//
type MetricType string

//goland:noinspection GoUnusedConst
const (
	// MetricTypeCounter is a data type that contains ever increasing numbers from the start of the server.
	MetricTypeCounter MetricType = "counter"
	// MetricTypeGauge is a metric type that can increase or decrease depending on the current value.
	MetricTypeGauge MetricType = "gauge"
)
