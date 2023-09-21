package healthcheckgrpcextention

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net/http"
)

type healthCheckGRPCExtension struct {
	config   Config
	logger   *zap.Logger
	server   *grpc.Server
	stopCh   chan struct{}
	settings component.TelemetrySettings
}

func (gc *healthCheckGRPCExtension) Start(_ context.Context, host component.Host) error {
	gc.logger.Info("Starting health_check_grpc extension", zap.Any("config", gc.config))
	ln, err := gc.config.ToListener()
	if err != nil {
		return fmt.Errorf("failed to bind to address %s: %w", gc.config.NetAddr.Endpoint, err)
	}

	gc.server, err = gc.config.ToServer(host, gc.settings)
	if err != nil {
		return err
	}

	gc.stopCh = make(chan struct{})
	hs := health.NewServer()

	// Register the health server with the gRPC server
	healthpb.RegisterHealthServer(gc.server, hs)
	reflection.Register(gc.server)

	go func() {
		defer close(gc.stopCh)

		// The listener ownership goes to the server.
		if err = gc.server.Serve(ln); !errors.Is(err, http.ErrServerClosed) && err != nil {
			host.ReportFatalError(err)
		}
	}()

	return nil
}

func (gc *healthCheckGRPCExtension) Shutdown(context.Context) error {
	if gc.server == nil {
		return nil
	}
	gc.server.GracefulStop()
	if gc.stopCh != nil {
		<-gc.stopCh
	}
	return nil
}

func newServer(config Config, settings component.TelemetrySettings) *healthCheckGRPCExtension {
	return &healthCheckGRPCExtension{
		config:   config,
		logger:   settings.Logger,
		settings: settings,
	}
}
