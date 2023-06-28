package config

import (
	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type IClientConfig interface {
	Connect() (onClose func(), err error)
	GetClient() (*grpc.ClientConn, error)
	GetFullyQualifiedDomainName() string
	GetIsSecure() bool
	GetTenantId() string
	GetDeviceId() string
}

// ClientConfig establishes the necessary parameters to establish a gRPC connection to Sensory Cloud
type ClientConfig struct {
	FullyQualifiedDomainName string
	IsSecure                 bool
	TenantId                 string
	DeviceId                 string
	client                   *grpc.ClientConn
}

// Connect establishes a gRPC connection to Sensory Cloud
func (c *ClientConfig) Connect() (onClose func(), err error) {
	client, onClose, err := grpc_client.NewCloudClient(c.FullyQualifiedDomainName, c.IsSecure)
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

// Gets the domain name
func (c *ClientConfig) GetFullyQualifiedDomainName() string {
	return c.FullyQualifiedDomainName
}

// Returns whether or not the client is secure
func (c *ClientConfig) GetIsSecure() bool {
	return c.IsSecure
}

// Gets the tenant id
func (c *ClientConfig) GetTenantId() string {
	return c.TenantId
}

// Gets the device id
func (c *ClientConfig) GetDeviceId() string {
	return c.DeviceId
}

// Enum to specify the authentication level required for device enrollment by the Sensory Cloud Server
type EnrollmentType int

const (
	// No authentication is required to enroll a new device
	None EnrollmentType = iota
	// Devices are enrolled via a shared secret (i.e. password)
	SharedSecret
	// Devices are enrolled via a signed JWT
	Jwt
)

// All configurations required to initialize the Sensory Cloud SDK
type SDKInitConfig struct {
	// The fully qualified domain name of teh Sensory Cloud Server to communicate with
	FullyQualifiedDomainName string
	// Tells if the SDK should use a secure connection to the Sensory Cloud server or not
	IsSecure bool
	// The tenant ID to use during device enrollment
	TenantID string
	// The level of authentication required to enroll new devices into the Sensory Cloud Server
	// If the device has already been enrolled during a previous session, this field is ignored
	EnrollmentType EnrollmentType
	// Credential for device enrollment
	// If the device has already been enrolled during a previous session, this field is ignored
	Credential string
	// Unique identifier for the current device
	DeviceID string
	// Friendly name for the current device
	DeviceName string
}
