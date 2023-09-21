package healthcheckgrpcextention

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckgrpcextension/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/extension"
)

const (
	defaultEndpoint = "0.0.0.0:13134"
)

// NewFactory creates a factory for HealthCheck extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		configgrpc.GRPCServerSettings{
			NetAddr: confignet.NetAddr{
				Endpoint:  defaultEndpoint,
				Transport: "tcp",
			},
		},
	}
}

func createExtension(_ context.Context, set extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	config := cfg.(*Config)

	return newServer(*config, set.TelemetrySettings), nil
}
