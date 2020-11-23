package metrics

import (
	"github.com/containerssh/http"
)

type Config struct {
	http.ServerConfiguration

	Enable bool   `yaml:"enable" json:"enable" comment:"Enable metrics server." default:"false"`
	Path   string `yaml:"path" json:"path" comment:"Path to run the Metrics endpoint on." default:"/metrics"`
}
