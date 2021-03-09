package metrics

import (
	"github.com/containerssh/http"
	"github.com/containerssh/log"
)

// NewServer creates a new metrics server based on the configuration. It MAY return nil if the server is disabled.
func NewServer(config Config, collector Collector, logger log.Logger) (http.Server, error) {
	if !config.Enable {
		return nil, nil
	}
	return http.NewServer(
		"Metrics server",
		config.ServerConfiguration,
		NewHandler(
			config.Path,
			collector,
		),
		logger,
		func(url string) {
			logger.Info(log.NewMessage(
				MServiceAvailable,
				"Metrics server is now available at %s%s",
				url, config.Path,
			))
		},
	)
}
