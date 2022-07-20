package initializer_test

import (
	"context"
	"errors"
	"testing"

	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	management_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/management"
	initializer "github.com/Sensory-Cloud/go-sdk/pkg/initializer"
	oauth_service "github.com/Sensory-Cloud/go-sdk/pkg/service/oauth"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

var (
	mockService         = NewMockOauthService()
	mockCredentialStore = NewMockSecureTokenStore()
	mockSdkConfig       = config.SDKInitConfig{
		FullyQualifiedDomainName: "domainName",
		IsSecure:                 true,
		TenantID:                 "tenant",
		EnrollmentType:           config.SharedSecret,
		Credential:               "credential",
		DeviceID:                 "deviceID",
		DeviceName:               "deviceName",
	}
	expectedClientConfig = config.ClientConfig{
		FullyQualifiedDomainName: "domainName",
		IsSecure:                 true,
		TenantId:                 "tenant",
		DeviceId:                 "deviceID",
	}
	expectedDeviceResponse = management_api_v1.DeviceResponse{Name: "deviceName", DeviceId: "deviceID"}
)

func TestInitialize(t *testing.T) {
	initializer.OauthServiceInit = GetMockOauthService
	mockService.Reset()
	mockCredentialStore.SaveCredentials("", "")
	mockService.Response = &expectedDeviceResponse

	context := context.Background()

	sdkInitializer := initializer.NewInitializer(mockCredentialStore, nil)
	clientConfig, rsp, err := sdkInitializer.Initialize(context, mockSdkConfig)

	// Assert responses
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, clientConfig, &expectedClientConfig)
	test_util.AssertEquals(t, rsp.Name, mockSdkConfig.DeviceName)
	test_util.AssertEquals(t, rsp.DeviceId, mockSdkConfig.DeviceID)
	// Asserts for secure credential store
	test_util.Assert(t, mockCredentialStore.clientId != "", "clientID should be saved")
	test_util.Assert(t, mockCredentialStore.secret != "", "client secret should be saved")
	// Asserts for OAuth service
	test_util.AssertEquals(t, mockService.DeviceName, mockSdkConfig.DeviceName)
	test_util.AssertEquals(t, mockService.Credential, mockSdkConfig.Credential)
}

func TestInitializeSavedCredentials(t *testing.T) {
	initializer.OauthServiceInit = GetMockOauthService
	mockService.Reset()
	mockCredentialStore.SaveCredentials("Saved-ID", "Saved-Secret")

	context := context.Background()

	sdkInitializer := initializer.NewInitializer(mockCredentialStore, nil)
	clientConfig, rsp, err := sdkInitializer.Initialize(context, mockSdkConfig)

	// Assert responses
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, clientConfig, &expectedClientConfig)
	test_util.Assert(t, rsp == nil, "No response should be returned when device has already been enrolled")
	// Assert Oauth service
	test_util.Assert(t, mockService.DeviceName == "", "Oauth service should not be called when device has already been enrolled")
}

// Mock Secure Credential Store
type mockSecureCredentialStore struct {
	clientId string
	secret   string
}

func NewMockSecureTokenStore() *mockSecureCredentialStore {
	return &mockSecureCredentialStore{"", ""}
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

// Mock Oauth service
type mockOauthService struct {
	Ctx        context.Context
	DeviceName string
	Credential string
	Response   *management_api_v1.DeviceResponse
	Err        error
}

func NewMockOauthService() *mockOauthService {
	return &mockOauthService{nil, "", "", nil, nil}
}

func GetMockOauthService(config *config.ClientConfig, credentialStore oauth_service.ISecureCredentialStore) (oauth_service.IOauthService, error) {
	return mockService, nil
}

func (m *mockOauthService) Reset() {
	m.Ctx = nil
	m.DeviceName = ""
	m.Credential = ""
	m.Response = nil
	m.Err = nil
}

func (m *mockOauthService) Register(ctx context.Context, deviceName string, credential string) (*management_api_v1.DeviceResponse, error) {
	m.Ctx = ctx
	m.DeviceName = deviceName
	m.Credential = credential
	return m.Response, m.Err
}

// Unimplemented mocks
func (m *mockOauthService) GetWhoAmI(ctx context.Context) (*management_api_v1.DeviceResponse, error) {
	return nil, errors.New("Unimplemented for this mock")
}

func (m *mockOauthService) GetToken(ctx context.Context) (*common.TokenResponse, error) {
	return nil, errors.New("Unimplemented for this mock")
}

func (m *mockOauthService) RenewDeviceCredential(ctx context.Context, credential string) (*management_api_v1.DeviceResponse, error) {
	return nil, errors.New("Unimplemented for this mock")
}
