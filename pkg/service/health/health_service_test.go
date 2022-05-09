package health_service_test

import (
	"context"
	"testing"

	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	health_service "github.com/Sensory-Cloud/go-sdk/pkg/service/health"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"github.com/google/uuid"
)

func TestGetHealth(t *testing.T) {
	config := config.ClientConfig{
		FullyQualifiedDomainName: "localhost:50050",
		TenantId:                 uuid.NewString(),
		DeviceId:                 uuid.NewString(),
	}

	onClose, err := config.Connect()
	test_util.AssertOk(t, err)
	defer onClose()

	healthService, err := health_service.NewHealthService(&config)
	test_util.AssertOk(t, err)

	response, err := healthService.GetHealth(context.Background())
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, response.IsHealthy, true)
}
