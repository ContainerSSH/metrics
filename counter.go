package metrics

type counterImpl struct {
	name      string
	collector *collector
}

func (c *counterImpl) Increment(labels ...MetricLabel) {
	_ = c.IncrementBy(1, labels...)
}

func (c *counterImpl) IncrementBy(by float64, labels ...MetricLabel) error {
	c.collector.mutex.Lock()
	defer c.collector.mutex.Unlock()

	if by < 0 {
		return CounterCannotBeIncrementedByNegative
	}

	realLabels := metricLabels(labels).toMap()
	value := c.collector.get(c.name, realLabels)
	c.collector.set(c.name, realLabels, value+by)
	return nil
}
