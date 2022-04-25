package grpc_client_test

import (
	"testing"

	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestGetDefaultContext(t *testing.T) {
	ctx, cancel := grpc_client.GetDefaultContext()
	test_util.Assert(t, ctx != nil, "context should not be nil")
	test_util.Assert(t, cancel != nil, "cancel should not be nil")
}
