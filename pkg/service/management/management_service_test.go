package management_service_test

import (
	"context"
	"testing"

	management_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/management"
	management_service "github.com/Sensory-Cloud/go-sdk/pkg/service/management"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestNewManagementService(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}

	managementService, err := management_service.NewManagementService(clientConfig, tokenManager)
	test_util.AssertOk(t, err)
	test_util.Assert(t, managementService != nil, "management service should not be nil")
}

func TestManagementService_GetEnrollments(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	enrollmentClient := &mockEnrollmentClient{}

	managementService := management_service.ManagementService{}
	managementService.SetConfig(clientConfig)
	managementService.SetTokenManager(tokenManager)
	managementService.SetEnrollmentClient(enrollmentClient)

	ctx := context.Background()

	response, err := managementService.GetEnrollments(ctx, "user-id")
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	managementService.SetTokenManager(&mockTokenManager{true})

	response, err = managementService.GetEnrollments(ctx, "user-id")
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
}

func TestManagementService_GetEnrollmentGroups(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	enrollmentClient := &mockEnrollmentClient{}

	managementService := management_service.ManagementService{}
	managementService.SetConfig(clientConfig)
	managementService.SetTokenManager(tokenManager)
	managementService.SetEnrollmentClient(enrollmentClient)

	ctx := context.Background()

	response, err := managementService.GetEnrollmentGroups(ctx, "user-id")
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	managementService.SetTokenManager(&mockTokenManager{true})

	response, err = managementService.GetEnrollmentGroups(ctx, "user-id")
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
}

func TestManagementService_CreateEnrollmentGroups(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	enrollmentClient := &mockEnrollmentClient{}

	managementService := management_service.ManagementService{}
	managementService.SetConfig(clientConfig)
	managementService.SetTokenManager(tokenManager)
	managementService.SetEnrollmentClient(enrollmentClient)

	ctx := context.Background()

	response, err := managementService.CreateEnrollmentGroups(ctx, &management_api_v1.CreateEnrollmentGroupRequest{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	managementService.SetTokenManager(&mockTokenManager{true})

	response, err = managementService.CreateEnrollmentGroups(ctx, &management_api_v1.CreateEnrollmentGroupRequest{})
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
}

func TestManagementService_AppendEnrollmentGroup(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	enrollmentClient := &mockEnrollmentClient{}

	managementService := management_service.ManagementService{}
	managementService.SetConfig(clientConfig)
	managementService.SetTokenManager(tokenManager)
	managementService.SetEnrollmentClient(enrollmentClient)

	ctx := context.Background()

	response, err := managementService.AppendEnrollmentGroup(ctx, &management_api_v1.AppendEnrollmentGroupRequest{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	managementService.SetTokenManager(&mockTokenManager{true})

	response, err = managementService.AppendEnrollmentGroup(ctx, &management_api_v1.AppendEnrollmentGroupRequest{})
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
}

func TestManagementService_DeleteEnrollment(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	enrollmentClient := &mockEnrollmentClient{}

	managementService := management_service.ManagementService{}
	managementService.SetConfig(clientConfig)
	managementService.SetTokenManager(tokenManager)
	managementService.SetEnrollmentClient(enrollmentClient)

	ctx := context.Background()

	response, err := managementService.DeleteEnrollment(ctx, "id")
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	managementService.SetTokenManager(&mockTokenManager{true})

	response, err = managementService.DeleteEnrollment(ctx, "id")
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
}

func TestManagementService_DeleteEnrollmentGroup(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	enrollmentClient := &mockEnrollmentClient{}

	managementService := management_service.ManagementService{}
	managementService.SetConfig(clientConfig)
	managementService.SetTokenManager(tokenManager)
	managementService.SetEnrollmentClient(enrollmentClient)

	ctx := context.Background()

	response, err := managementService.DeleteEnrollmentGroup(ctx, "id")
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	managementService.SetTokenManager(&mockTokenManager{true})

	response, err = managementService.DeleteEnrollmentGroup(ctx, "id")
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
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

// Mock Token Manager
type mockTokenManager struct {
	AuthorizationMetadataError bool
}

func (c *mockTokenManager) SetAuthorizationMetadata(ctx context.Context) (context.Context, error) {
	if c.AuthorizationMetadataError {
		return nil, error_util.Error(codes.InvalidArgument, "set authorization metadata test error")
	}
	return nil, nil
}

func (c *mockTokenManager) GetToken() (string, error) {
	return "", nil
}

func (c *mockTokenManager) RefreshToken() error {
	return nil
}

func (c *mockTokenManager) IsTokenExpired() bool {
	return false
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

// Mock Enrollment Client
type mockEnrollmentClient struct{}

func (c *mockEnrollmentClient) GetEnrollments(ctx context.Context, in *management_api_v1.GetEnrollmentsRequest, opts ...grpc.CallOption) (*management_api_v1.GetEnrollmentsResponse, error) {
	return &management_api_v1.GetEnrollmentsResponse{}, nil
}
func (c *mockEnrollmentClient) GetEnrollmentGroups(ctx context.Context, in *management_api_v1.GetEnrollmentsRequest, opts ...grpc.CallOption) (*management_api_v1.GetEnrollmentGroupsResponse, error) {
	return &management_api_v1.GetEnrollmentGroupsResponse{}, nil
}
func (c *mockEnrollmentClient) CreateEnrollmentGroup(ctx context.Context, in *management_api_v1.CreateEnrollmentGroupRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentGroupResponse, error) {
	return &management_api_v1.EnrollmentGroupResponse{}, nil
}
func (c *mockEnrollmentClient) AppendEnrollmentGroup(ctx context.Context, in *management_api_v1.AppendEnrollmentGroupRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentGroupResponse, error) {
	return &management_api_v1.EnrollmentGroupResponse{}, nil
}
func (c *mockEnrollmentClient) DeleteEnrollment(ctx context.Context, in *management_api_v1.DeleteEnrollmentRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentResponse, error) {
	return &management_api_v1.EnrollmentResponse{}, nil
}
func (c *mockEnrollmentClient) DeleteEnrollmentGroup(ctx context.Context, in *management_api_v1.DeleteEnrollmentGroupRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentGroupResponse, error) {
	return &management_api_v1.EnrollmentGroupResponse{}, nil
}
func (c *mockEnrollmentClient) UpdateEnrollment(ctx context.Context, in *management_api_v1.UpdateEnrollmentRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentResponse, error) {
	return &management_api_v1.EnrollmentResponse{}, nil
}
func (c *mockEnrollmentClient) UpdateEnrollmentGroup(ctx context.Context, in *management_api_v1.UpdateEnrollmentGroupRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentGroupResponse, error) {
	return &management_api_v1.EnrollmentGroupResponse{}, nil
}
func (c *mockEnrollmentClient) RemoveEnrollmentsFromGroup(ctx context.Context, in *management_api_v1.RemoveEnrollmentsRequest, opts ...grpc.CallOption) (*management_api_v1.EnrollmentGroupResponse, error) {
	return &management_api_v1.EnrollmentGroupResponse{}, nil
}
