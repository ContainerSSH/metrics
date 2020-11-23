package metrics

type counterImpl struct {
	name      string
	collector *collector
}

func (c *counterImpl) Increment() {
	_ = c.IncrementBy(1)
}

func (c *counterImpl) IncrementBy(by float64) error {
	c.collector.mutex.Lock()
	defer c.collector.mutex.Unlock()

	if by < 0 {
		return CounterCannotBeIncrementedByNegative
	}

	value := c.collector.get(c.name, map[string]string{})
	c.collector.set(c.name, map[string]string{}, value+by)
	return nil
}
