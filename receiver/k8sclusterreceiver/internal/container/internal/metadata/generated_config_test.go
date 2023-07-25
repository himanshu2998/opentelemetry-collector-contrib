// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sContainerCPULimit:                MetricConfig{Enabled: true},
					K8sContainerCPURequest:              MetricConfig{Enabled: true},
					K8sContainerEphemeralstorageLimit:   MetricConfig{Enabled: true},
					K8sContainerEphemeralstorageRequest: MetricConfig{Enabled: true},
					K8sContainerMemoryLimit:             MetricConfig{Enabled: true},
					K8sContainerMemoryRequest:           MetricConfig{Enabled: true},
					K8sContainerReady:                   MetricConfig{Enabled: true},
					K8sContainerRestarts:                MetricConfig{Enabled: true},
					K8sContainerStorageLimit:            MetricConfig{Enabled: true},
					K8sContainerStorageRequest:          MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerID:            ResourceAttributeConfig{Enabled: true},
					ContainerImageName:     ResourceAttributeConfig{Enabled: true},
					ContainerImageTag:      ResourceAttributeConfig{Enabled: true},
					K8sContainerName:       ResourceAttributeConfig{Enabled: true},
					K8sNamespaceName:       ResourceAttributeConfig{Enabled: true},
					K8sNodeName:            ResourceAttributeConfig{Enabled: true},
					K8sPodName:             ResourceAttributeConfig{Enabled: true},
					K8sPodUID:              ResourceAttributeConfig{Enabled: true},
					OpencensusResourcetype: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sContainerCPULimit:                MetricConfig{Enabled: false},
					K8sContainerCPURequest:              MetricConfig{Enabled: false},
					K8sContainerEphemeralstorageLimit:   MetricConfig{Enabled: false},
					K8sContainerEphemeralstorageRequest: MetricConfig{Enabled: false},
					K8sContainerMemoryLimit:             MetricConfig{Enabled: false},
					K8sContainerMemoryRequest:           MetricConfig{Enabled: false},
					K8sContainerReady:                   MetricConfig{Enabled: false},
					K8sContainerRestarts:                MetricConfig{Enabled: false},
					K8sContainerStorageLimit:            MetricConfig{Enabled: false},
					K8sContainerStorageRequest:          MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerID:            ResourceAttributeConfig{Enabled: false},
					ContainerImageName:     ResourceAttributeConfig{Enabled: false},
					ContainerImageTag:      ResourceAttributeConfig{Enabled: false},
					K8sContainerName:       ResourceAttributeConfig{Enabled: false},
					K8sNamespaceName:       ResourceAttributeConfig{Enabled: false},
					K8sNodeName:            ResourceAttributeConfig{Enabled: false},
					K8sPodName:             ResourceAttributeConfig{Enabled: false},
					K8sPodUID:              ResourceAttributeConfig{Enabled: false},
					OpencensusResourcetype: ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				ContainerID:            ResourceAttributeConfig{Enabled: true},
				ContainerImageName:     ResourceAttributeConfig{Enabled: true},
				ContainerImageTag:      ResourceAttributeConfig{Enabled: true},
				K8sContainerName:       ResourceAttributeConfig{Enabled: true},
				K8sNamespaceName:       ResourceAttributeConfig{Enabled: true},
				K8sNodeName:            ResourceAttributeConfig{Enabled: true},
				K8sPodName:             ResourceAttributeConfig{Enabled: true},
				K8sPodUID:              ResourceAttributeConfig{Enabled: true},
				OpencensusResourcetype: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				ContainerID:            ResourceAttributeConfig{Enabled: false},
				ContainerImageName:     ResourceAttributeConfig{Enabled: false},
				ContainerImageTag:      ResourceAttributeConfig{Enabled: false},
				K8sContainerName:       ResourceAttributeConfig{Enabled: false},
				K8sNamespaceName:       ResourceAttributeConfig{Enabled: false},
				K8sNodeName:            ResourceAttributeConfig{Enabled: false},
				K8sPodName:             ResourceAttributeConfig{Enabled: false},
				K8sPodUID:              ResourceAttributeConfig{Enabled: false},
				OpencensusResourcetype: ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
