package oauth_service_test

import (
	"context"
	"testing"

	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	oauth_api "github.com/Sensory-Cloud/go-sdk/pkg/generated/oauth"
	management_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/management"

	oauth_service "github.com/Sensory-Cloud/go-sdk/pkg/service/oauth"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"google.golang.org/grpc"
)

func TestNewOauthService(t *testing.T) {
	credentialStore := &mockSecureCredentialStore{"client-id", "secret"}
	clientConfig := &mockClientConfig{}

	oauthService, err := oauth_service.NewOauthService(clientConfig, credentialStore)
	test_util.AssertOk(t, err)
	test_util.Assert(t, oauthService != nil, "oauth service should not be nil")
}

func TestGenerateCredentials(t *testing.T) {
	oauthClient, err := oauth_service.GenerateCredentials()
	test_util.AssertOk(t, err)
	test_util.Assert(t, oauthClient.ClientId != "", "client id should not be an empty string")
	test_util.Assert(t, oauthClient.ClientSecret != "", "client secret should not be an empty string")
}

func TestGetWhoAmI(t *testing.T) {
	credentialStore := &mockSecureCredentialStore{"client-id", "secret"}
	clientConfig := &mockClientConfig{}
	oauthClient := &mockOauthServiceClient{}
	deviceClient := &mockDeviceServiceClient{}

	oauthService := &oauth_service.OauthService{}
	oauthService.SetConfig(clientConfig)
	oauthService.SetSecureCredentialStore(credentialStore)
	oauthService.SetOauthClient(oauthClient)
	oauthService.SetDeviceClient(deviceClient)

	ctx := context.Background()

	response, err := oauthService.GetWhoAmI(ctx)
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")
}

func TestGetToken(t *testing.T) {
	credentialStore := &mockSecureCredentialStore{"client-id", "secret"}
	clientConfig := &mockClientConfig{}
	oauthClient := &mockOauthServiceClient{}

	oauthService := &oauth_service.OauthService{}
	oauthService.SetConfig(clientConfig)
	oauthService.SetSecureCredentialStore(credentialStore)
	oauthService.SetOauthClient(oauthClient)

	ctx := context.Background()

	response, err := oauthService.GetToken(ctx)
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")
}

func TestRegister(t *testing.T) {
	credentialStore := &mockSecureCredentialStore{"client-id", "secret"}
	clientConfig := &mockClientConfig{}
	deviceClient := &mockDeviceServiceClient{}

	oauthService := &oauth_service.OauthService{}
	oauthService.SetConfig(clientConfig)
	oauthService.SetSecureCredentialStore(credentialStore)
	oauthService.SetDeviceClient(deviceClient)

	ctx := context.Background()

	response, err := oauthService.Register(ctx, "device-name", "credential")
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")
}

func TestRenewDeviceCredential(t *testing.T) {
	credentialStore := &mockSecureCredentialStore{"client-id", "secret"}
	clientConfig := &mockClientConfig{}
	deviceClient := &mockDeviceServiceClient{}

	oauthService := &oauth_service.OauthService{}
	oauthService.SetConfig(clientConfig)
	oauthService.SetSecureCredentialStore(credentialStore)
	oauthService.SetDeviceClient(deviceClient)

	ctx := context.Background()

	response, err := oauthService.RenewDeviceCredential(ctx, "credential")
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")
}

// Mock Secure Credential Store
type mockSecureCredentialStore struct {
	clientId string
	secret   string
}

func (m *mockSecureCredentialStore) SaveCredentials(clientId string, secret string) {
	m.clientId = clientId
	m.secret = secret
}

func (m *mockSecureCredentialStore) GetClientId() string {
	return m.clientId
}

func (m *mockSecureCredentialStore) GetClientSecret() string {
	return m.secret
}

// Mock Client Config
type mockClientConfig struct {
	FullyQualifiedDomainName string
	IsSecure                 bool
	TenantId                 string
	DeviceId                 string
}

func (c *mockClientConfig) Connect() (onClose func(), err error) {
	return func() {}, nil
}

func (c *mockClientConfig) GetClient() (*grpc.ClientConn, error) {
	return &grpc.ClientConn{}, nil
}

func (c *mockClientConfig) GetFullyQualifiedDomainName() string {
	return c.FullyQualifiedDomainName
}

func (c *mockClientConfig) GetIsSecure() bool {
	return c.IsSecure
}

func (c *mockClientConfig) GetTenantId() string {
	return c.TenantId
}

func (c *mockClientConfig) GetDeviceId() string {
	return c.DeviceId
}

// Mock Oauth Service Client
type mockOauthServiceClient struct{}

func (c *mockOauthServiceClient) GetToken(ctx context.Context, in *oauth_api.TokenRequest, opts ...grpc.CallOption) (*common.TokenResponse, error) {
	return &common.TokenResponse{}, nil
}
func (c *mockOauthServiceClient) SignToken(ctx context.Context, in *oauth_api.SignTokenRequest, opts ...grpc.CallOption) (*common.TokenResponse, error) {
	return nil, nil
}
func (c *mockOauthServiceClient) GetWhoAmI(ctx context.Context, in *oauth_api.WhoAmIRequest, opts ...grpc.CallOption) (*oauth_api.WhoAmIResponse, error) {
	return nil, nil
}
func (c *mockOauthServiceClient) GetPublicKey(ctx context.Context, in *oauth_api.PublicKeyRequest, opts ...grpc.CallOption) (*oauth_api.PublicKeyResponse, error) {
	return nil, nil
}

// Mock Device Service Client
type mockDeviceServiceClient struct{}

func (c *mockDeviceServiceClient) EnrollDevice(ctx context.Context, in *management_api_v1.EnrollDeviceRequest, opts ...grpc.CallOption) (*management_api_v1.DeviceResponse, error) {
	return &management_api_v1.DeviceResponse{}, nil
}
func (c *mockDeviceServiceClient) RenewDeviceCredential(ctx context.Context, in *management_api_v1.RenewDeviceCredentialRequest, opts ...grpc.CallOption) (*management_api_v1.DeviceResponse, error) {
	return &management_api_v1.DeviceResponse{}, nil
}
func (c *mockDeviceServiceClient) GetWhoAmI(ctx context.Context, in *management_api_v1.DeviceGetWhoAmIRequest, opts ...grpc.CallOption) (*management_api_v1.DeviceResponse, error) {
	return &management_api_v1.DeviceResponse{}, nil
}
func (c *mockDeviceServiceClient) GetDevice(ctx context.Context, in *management_api_v1.DeviceRequest, opts ...grpc.CallOption) (*management_api_v1.GetDeviceResponse, error) {
	return nil, nil
}
func (c *mockDeviceServiceClient) GetDevices(ctx context.Context, in *management_api_v1.GetDevicesRequest, opts ...grpc.CallOption) (*management_api_v1.DeviceListResponse, error) {
	return nil, nil
}
func (c *mockDeviceServiceClient) UpdateDevice(ctx context.Context, in *management_api_v1.UpdateDeviceRequest, opts ...grpc.CallOption) (*management_api_v1.DeviceResponse, error) {
	return nil, nil
}
func (c *mockDeviceServiceClient) DeleteDevice(ctx context.Context, in *management_api_v1.DeviceRequest, opts ...grpc.CallOption) (*management_api_v1.DeviceResponse, error) {
	return nil, nil
}
