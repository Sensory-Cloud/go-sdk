package video_service_test

import (
	"context"
	"testing"

	video_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/video"
	video_service "github.com/Sensory-Cloud/go-sdk/pkg/service/video"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func TestNewVideoService(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}

	videoService, err := video_service.NewVideoService(clientConfig, tokenManager)
	test_util.AssertOk(t, err)
	test_util.Assert(t, videoService != nil, "video service should not be nil")
}

func TestGetModels(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	videoModelsClient := &mockVideoModelsClient{}

	videoService := video_service.VideoService{}
	videoService.SetClientConfig(clientConfig)
	videoService.SetTokenManager(tokenManager)
	videoService.SetVideoModelsClient(videoModelsClient)

	ctx := context.Background()

	response, err := videoService.GetModels(ctx)
	test_util.AssertOk(t, err)
	test_util.Assert(t, response != nil, "response should not be nil")

	videoService.SetTokenManager(&mockTokenManager{true})

	response, err = videoService.GetModels(ctx)
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, response == nil, "response should be nil")
}

func TestStreamEnrollment(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	videoBiometricsClient := &mockVideoBiometricsClient{}

	videoService := video_service.VideoService{}
	videoService.SetClientConfig(clientConfig)
	videoService.SetTokenManager(tokenManager)
	videoService.SetVideoBiometricsClient(videoBiometricsClient)

	ctx := context.Background()

	stream, err := videoService.StreamEnrollment(ctx, &video_api_v1.CreateEnrollmentConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	videoService.SetTokenManager(&mockTokenManager{true})

	stream, err = videoService.StreamEnrollment(ctx, &video_api_v1.CreateEnrollmentConfig{})
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestStreamAuthentication(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	videoBiometricsClient := &mockVideoBiometricsClient{}

	videoService := video_service.VideoService{}
	videoService.SetClientConfig(clientConfig)
	videoService.SetTokenManager(tokenManager)
	videoService.SetVideoBiometricsClient(videoBiometricsClient)

	ctx := context.Background()

	stream, err := videoService.StreamAuthentication(ctx, &video_api_v1.AuthenticateConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	videoService.SetTokenManager(&mockTokenManager{true})

	stream, err = videoService.StreamAuthentication(ctx, &video_api_v1.AuthenticateConfig{})
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestStreamLivenessRecognition(t *testing.T) {
	clientConfig := &mockClientConfig{}
	tokenManager := &mockTokenManager{}
	videoRecognitionClient := &mockVideoRecognitionClient{}

	videoService := video_service.VideoService{}
	videoService.SetClientConfig(clientConfig)
	videoService.SetTokenManager(tokenManager)
	videoService.SetVideoRecognitionClient(videoRecognitionClient)

	ctx := context.Background()

	stream, err := videoService.StreamLivenessRecognition(ctx, &video_api_v1.ValidateRecognitionConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	videoService.SetTokenManager(&mockTokenManager{true})

	stream, err = videoService.StreamLivenessRecognition(ctx, &video_api_v1.ValidateRecognitionConfig{})
	test_util.Assert(t, err != nil, "error should not be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
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

// Mock Video Biometrics Client
type mockVideoBiometricsClient struct{}

func (c *mockVideoBiometricsClient) CreateEnrollment(ctx context.Context, opts ...grpc.CallOption) (video_api_v1.VideoBiometrics_CreateEnrollmentClient, error) {
	return &mockBiometricsCreateEnrollmentClient{}, nil
}
func (c *mockVideoBiometricsClient) Authenticate(ctx context.Context, opts ...grpc.CallOption) (video_api_v1.VideoBiometrics_AuthenticateClient, error) {
	return &mockBiometricsAuthenticateClient{}, nil
}

// Mock Biometrics Create Enrollment Client
type mockBiometricsCreateEnrollmentClient struct{}

func (c *mockBiometricsCreateEnrollmentClient) Send(m *video_api_v1.CreateEnrollmentRequest) error {
	return c.SendMsg(m)
}

func (c *mockBiometricsCreateEnrollmentClient) Recv() (*video_api_v1.CreateEnrollmentResponse, error) {
	m := new(video_api_v1.CreateEnrollmentResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockBiometricsCreateEnrollmentClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockBiometricsCreateEnrollmentClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockBiometricsCreateEnrollmentClient) CloseSend() error {
	return nil
}

func (stream *mockBiometricsCreateEnrollmentClient) Context() context.Context {
	return context.Background()
}

func (stream *mockBiometricsCreateEnrollmentClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockBiometricsCreateEnrollmentClient) RecvMsg(m interface{}) error {
	return nil
}

// Mock Biometrics Create Enrollment Client
type mockBiometricsAuthenticateClient struct{}

func (c *mockBiometricsAuthenticateClient) Send(m *video_api_v1.AuthenticateRequest) error {
	return c.SendMsg(m)
}

func (c *mockBiometricsAuthenticateClient) Recv() (*video_api_v1.AuthenticateResponse, error) {
	m := new(video_api_v1.AuthenticateResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockBiometricsAuthenticateClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockBiometricsAuthenticateClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockBiometricsAuthenticateClient) CloseSend() error {
	return nil
}

func (stream *mockBiometricsAuthenticateClient) Context() context.Context {
	return context.Background()
}

func (stream *mockBiometricsAuthenticateClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockBiometricsAuthenticateClient) RecvMsg(m interface{}) error {
	return nil
}

// Mock Video Models Client
type mockVideoModelsClient struct{}

func (c *mockVideoModelsClient) GetModels(ctx context.Context, in *video_api_v1.GetModelsRequest, opts ...grpc.CallOption) (*video_api_v1.GetModelsResponse, error) {
	return &video_api_v1.GetModelsResponse{}, nil
}

// Mock Video Recognition Client
type mockVideoRecognitionClient struct{}

func (c *mockVideoRecognitionClient) ValidateLiveness(ctx context.Context, opts ...grpc.CallOption) (video_api_v1.VideoRecognition_ValidateLivenessClient, error) {
	return &mockValidateLivenessClient{}, nil
}

// Mock Validate Liveness Client
type mockValidateLivenessClient struct{}

func (c *mockValidateLivenessClient) Send(m *video_api_v1.ValidateRecognitionRequest) error {
	return c.SendMsg(m)
}

func (c *mockValidateLivenessClient) Recv() (*video_api_v1.LivenessRecognitionResponse, error) {
	m := new(video_api_v1.LivenessRecognitionResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockValidateLivenessClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockValidateLivenessClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockValidateLivenessClient) CloseSend() error {
	return nil
}

func (stream *mockValidateLivenessClient) Context() context.Context {
	return context.Background()
}

func (stream *mockValidateLivenessClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockValidateLivenessClient) RecvMsg(m interface{}) error {
	return nil
}
