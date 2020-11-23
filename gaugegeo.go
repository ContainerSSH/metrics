package metrics

import (
	"net"
)

type gaugeGeoImpl struct {
	name      string
	collector *collector
}

func (g *gaugeGeoImpl) Increment(ip net.IP) {
	g.IncrementBy(ip, 1)
}

func (g *gaugeGeoImpl) IncrementBy(ip net.IP, by float64) {
	g.collector.mutex.Lock()
	defer g.collector.mutex.Unlock()

	labels := map[string]string{
		"country": g.collector.geoIpLookupProvider.Lookup(ip),
	}

	value := g.collector.get(g.name, labels)
	g.collector.set(g.name, labels, value+by)
}

func (g *gaugeGeoImpl) Decrement(ip net.IP) {
	g.DecrementBy(ip, 1)
}

func (g *gaugeGeoImpl) DecrementBy(ip net.IP, by float64) {
	g.collector.mutex.Lock()
	defer g.collector.mutex.Unlock()

	labels := map[string]string{
		"country": g.collector.geoIpLookupProvider.Lookup(ip),
	}

	value := g.collector.get(g.name, labels)
	g.collector.set(g.name, labels, value-by)
}

func (g *gaugeGeoImpl) Set(ip net.IP, value float64) {
	g.collector.mutex.Lock()
	defer g.collector.mutex.Unlock()

	labels := map[string]string{
		"country": g.collector.geoIpLookupProvider.Lookup(ip),
	}

	g.collector.set(g.name, labels, value)
}
