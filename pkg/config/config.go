package config

import (
	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ClientConfig establishes the necessary parameters to establish a gRPC connection to Sensory Cloud
type ClientConfig struct {
	FullyQualifiedDomainName string
	TenantId                 string
	DeviceId                 string
	client                   *grpc.ClientConn
}

// Connect establishes a gRPC connection to Sensory Cloud
func (c *ClientConfig) Connect() (onClose func(), err error) {
	client, onClose, err := grpc_client.NewCloudClient(c.FullyQualifiedDomainName)
	if err != nil {
		return onClose, err
	}

	c.client = client
	return onClose, nil
}

// GetClient returns the connected client. An error will be returned if Connect() has not been called
func (c *ClientConfig) GetClient() (*grpc.ClientConn, error) {
	if c.client == nil {
		return nil, error_util.Error(codes.FailedPrecondition, "no connection has been established with Sensory Cloud. did you forget to call Connect()?")
	}

	return c.client, nil
}
