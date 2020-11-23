package metrics

import (
	"net"
)

type counterGeoImpl struct {
	name      string
	collector *collector
}

func (c *counterGeoImpl) Increment(ip net.IP) {
	_ = c.IncrementBy(ip, 1)
}

func (c *counterGeoImpl) IncrementBy(ip net.IP, by float64) error {
	c.collector.mutex.Lock()
	defer c.collector.mutex.Unlock()

	if by < 0 {
		return CounterCannotBeIncrementedByNegative
	}

	labels := map[string]string{
		"country": c.collector.geoIpLookupProvider.Lookup(ip),
	}

	value := c.collector.get(c.name, labels)
	c.collector.set(c.name, labels, value+by)
	return nil
}
