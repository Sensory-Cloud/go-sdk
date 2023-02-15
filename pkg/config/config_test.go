package config_test

import (
	"testing"

	"github.com/Sensory-Cloud/go-sdk/pkg/config"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestConnect(t *testing.T) {
	clientConfig := config.ClientConfig{
		FullyQualifiedDomainName: "domain-name",
		IsSecure:                 true,
	}

	onClose, err := clientConfig.Connect()
	test_util.AssertOk(t, err)
	test_util.Assert(t, onClose != nil, "on close function should not be nil")
}

func TestGetClient(t *testing.T) {
	clientConfig := config.ClientConfig{
		FullyQualifiedDomainName: "domain-name",
		IsSecure:                 true,
	}

	client, err := clientConfig.GetClient()
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, client == nil, "client should be nil")

	_, err = clientConfig.Connect()
	test_util.AssertOk(t, err)

	client, err = clientConfig.GetClient()
	test_util.AssertOk(t, err)
	test_util.Assert(t, client != nil, "client should not be nil")
}

func TestGetAttributes(t *testing.T) {
	clientConfig := config.ClientConfig{
		FullyQualifiedDomainName: "domain-name",
		IsSecure:                 true,
		TenantId:                 "tenant-id",
		DeviceId:                 "device-id",
	}

	domainName := clientConfig.GetFullyQualifiedDomainName()
	test_util.AssertEquals(t, domainName, "domain-name")

	isSecure := clientConfig.GetIsSecure()
	test_util.AssertEquals(t, isSecure, true)

	tenantId := clientConfig.GetTenantId()
	test_util.AssertEquals(t, tenantId, "tenant-id")

	deviceId := clientConfig.GetDeviceId()
	test_util.AssertEquals(t, deviceId, "device-id")
}
