package metrics

import (
	"fmt"

	"github.com/containerssh/http"
)

type Config struct {
	http.ServerConfiguration `json:",inline" yaml:",inline" default:"{\"listen\":\"0.0.0.0:9100\"}"`

	Enable bool   `yaml:"enable" json:"enable" comment:"Enable metrics server." default:"false"`
	Path   string `yaml:"path" json:"path" comment:"Path to run the Metrics endpoint on." default:"/metrics"`
}

// Validate validates the configuration.
func (c Config) Validate() error {
	if !c.Enable {
		return nil
	}
	if c.Path == "" {
		return fmt.Errorf("metrics path cannot be empty")
	}
	return c.ServerConfiguration.Validate()
}
