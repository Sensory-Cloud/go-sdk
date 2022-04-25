package health_service

import (
	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	health_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/health"
)

type HealthService struct {
	config       *config.ClientConfig
	healthClient health_api_v1.HealthServiceClient
}

// NewHealthService creates a new server to check Sensory Cloud server health
func NewHealthService(config *config.ClientConfig) (*HealthService, error) {
	client, err := config.GetClient()
	if err != nil {
		return nil, err
	}

	healthClient := health_api_v1.NewHealthServiceClient(client)
	return &HealthService{config, healthClient}, nil
}

// GetHealth returns the current health status of the configured server
func (s *HealthService) GetHealth() (*common.ServerHealthResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	return s.healthClient.GetHealth(ctx, &health_api_v1.HealthRequest{})
}
