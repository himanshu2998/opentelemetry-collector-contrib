package healthcheckgrpcextention

import (
	"errors"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
)

type Config struct {
	configgrpc.GRPCServerSettings `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)
var (
	errNoEndpointProvided = errors.New("bad config: endpoint must be specified")
)

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error {
	if cfg.NetAddr.Endpoint == "" {
		return errNoEndpointProvided
	}
	return nil
}
