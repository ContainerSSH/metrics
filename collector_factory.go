package metrics

import (
	"sync"

	"github.com/containerssh/geoip"
)

// New creates the metric collector.
func New(geoIpLookupProvider geoip.LookupProvider) Collector {
	return &collector{
		geoIpLookupProvider: geoIpLookupProvider,
		mutex:               &sync.Mutex{},
		metricsMap:          map[string]Metric{},
		metrics:             []Metric{},
		values:              map[string]*metricValue{},
	}
}
