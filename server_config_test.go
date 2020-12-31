package metrics_test

import (
	"testing"

	"github.com/containerssh/structutils"
	"github.com/stretchr/testify/assert"

	"github.com/containerssh/metrics"
)

func TestListenDefault(t *testing.T) {
	config := metrics.Config{}
	structutils.Defaults(&config)
	assert.Equal(t, "0.0.0.0:9100", config.Listen)
}
