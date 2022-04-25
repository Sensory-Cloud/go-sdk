package grpc_client_test

import (
	"testing"

	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	environment "github.com/Sensory-Cloud/go-sdk/pkg/environment"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestGetDefaultContext(t *testing.T) {
	ctx, cancel := grpc_client.GetDefaultContext()
	test_util.Assert(t, ctx != nil, "context should not be nil")
	test_util.Assert(t, cancel != nil, "cancel should not be nil")
}

func TestNewCloudClientFromEnv(t *testing.T) {
	cloudHost := environment.Get(environment.CloudHost, "")
	defer environment.Set(environment.CloudHost, cloudHost)

	environment.Set(environment.CloudHost, "")
	cloudClient, cleanup, err := grpc_client.NewCloudClientFromEnv()
	defer cleanup()
	test_util.Assert(t, cloudClient == nil, "a cloud client should not be returned")
	test_util.Assert(t, err != nil, "an error should be returned")

	environment.Set(environment.CloudHost, "localhost:50050")
	grpc_client.NewCloudClientFromEnv()
}
