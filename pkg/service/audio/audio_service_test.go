package audio_service_test

import (
	"context"
	"testing"

	"github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	audio_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/audio"
	audio_service "github.com/Sensory-Cloud/go-sdk/pkg/service/audio"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func TestNewAudioService(t *testing.T) {
	audioService, err := audio_service.NewAudioService(&mockClientConfig{}, &mockTokenManager{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, audioService != nil, "audio service should not be nil")
}

func TestStreamEnrollment(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioBiometricsClient(&mockAudioBiometricsClient{})

	ctx := context.Background()
	stream, err := mockAudioService.StreamEnrollment(ctx, &audio_api_v1.CreateEnrollmentConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})

	stream, err = mockAudioService.StreamEnrollment(ctx, &audio_api_v1.CreateEnrollmentConfig{})
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestGetModels(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioModelsClient(&mockAudioModelsClient{})

	ctx := context.Background()
	models, err := mockAudioService.GetModels(ctx)
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, len(models.Models), 1)
	test_util.AssertEquals(t, models.Models[0].Name, "test-model")
	test_util.AssertEquals(t, models.Models[0].IsEnrollable, false)
	test_util.AssertEquals(t, models.Models[0].ModelType, common.ModelType_VOICE_BIOMETRIC_WAKEWORD)
	test_util.AssertEquals(t, models.Models[0].FixedPhrase, "fixed phrase")
	test_util.AssertEquals(t, models.Models[0].SampleRate, int32(16000))
	test_util.AssertEquals(t, models.Models[0].Versions, []string{"version0"})
	test_util.AssertEquals(t, models.Models[0].Technology, common.TechnologyType_NOT_SET)
	test_util.AssertEquals(t, models.Models[0].IsLivenessSupported, false)

	mockAudioService.SetTokenManager(&mockTokenManager{true})

	models, err = mockAudioService.GetModels(ctx)
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, models == nil, "stream should be nil")
}

func TestStreamAuthentication(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioBiometricsClient(&mockAudioBiometricsClient{})

	ctx := context.Background()
	stream, err := mockAudioService.StreamAuthentication(ctx, &audio_api_v1.AuthenticateConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})

	stream, err = mockAudioService.StreamAuthentication(ctx, &audio_api_v1.AuthenticateConfig{})
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestStreamCreateEnrolledEvent(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioEventsClient(&mockEventsClient{})

	ctx := context.Background()
	stream, err := mockAudioService.StreamCreateEnrolledEvent(ctx, &audio_api_v1.CreateEnrollmentEventConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})
	test_util.AssertOk(t, err)

	stream, err = mockAudioService.StreamCreateEnrolledEvent(ctx, &audio_api_v1.CreateEnrollmentEventConfig{})
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestStreamEvent(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioEventsClient(&mockEventsClient{})

	ctx := context.Background()
	stream, err := mockAudioService.StreamEvent(ctx, &audio_api_v1.ValidateEventConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})

	stream, err = mockAudioService.StreamEvent(ctx, &audio_api_v1.ValidateEventConfig{})
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestStreamValidateEnrolledEvent(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioEventsClient(&mockEventsClient{})

	ctx := context.Background()
	stream, err := mockAudioService.StreamValidateEnrolledEvent(ctx, &audio_api_v1.ValidateEnrolledEventConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})

	stream, err = mockAudioService.StreamValidateEnrolledEvent(ctx, &audio_api_v1.ValidateEnrolledEventConfig{})
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestStreamTranscription(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioTranscriptionsClient(&mockAudioTranscriptionClient{})

	ctx := context.Background()
	stream, err := mockAudioService.StreamTranscription(ctx, &audio_api_v1.TranscribeConfig{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})
	test_util.AssertOk(t, err)

	stream, err = mockAudioService.StreamTranscription(ctx, &audio_api_v1.TranscribeConfig{})
	test_util.Assert(t, err != nil, "err should be nil")
	test_util.Assert(t, stream == nil, "stream should be nil")
}

func TestSynthesizeSpeech(t *testing.T) {
	mockAudioService := audio_service.AudioService{}
	mockAudioService.SetConfig(&mockClientConfig{})
	mockAudioService.SetTokenManager(&mockTokenManager{})
	mockAudioService.SetAudioSynthesisClient(&mockAudioSynthesisClient{})

	ctx := context.Background()
	stream, err := mockAudioService.SynthesizeSpeech(ctx, &audio_api_v1.SynthesizeSpeechRequest{})
	test_util.AssertOk(t, err)
	test_util.Assert(t, stream != nil, "stream should not be nil")

	mockAudioService.SetTokenManager(&mockTokenManager{true})

	stream, err = mockAudioService.SynthesizeSpeech(ctx, &audio_api_v1.SynthesizeSpeechRequest{})
	test_util.Assert(t, err != nil, "err should be nil")
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

// MockAudioBiometricsClient
type mockAudioBiometricsClient struct {
}

func (c *mockAudioBiometricsClient) CreateEnrollment(ctx context.Context, opts ...grpc.CallOption) (audio_api_v1.AudioBiometrics_CreateEnrollmentClient, error) {
	biometricsEnrollmentClient := &mockBiometricsCreateEnrollmentClient{}
	return biometricsEnrollmentClient, nil
}

func (c *mockAudioBiometricsClient) Authenticate(ctx context.Context, opts ...grpc.CallOption) (audio_api_v1.AudioBiometrics_AuthenticateClient, error) {
	biometricsAuthenticateClient := &mockBiometricsAuthenticateClient{}
	return biometricsAuthenticateClient, nil
}

// Mock Create Enrollment Client
type mockBiometricsCreateEnrollmentClient struct{}

func (c *mockBiometricsCreateEnrollmentClient) Send(m *audio_api_v1.CreateEnrollmentRequest) error {
	return c.SendMsg(m)
}

func (c *mockBiometricsCreateEnrollmentClient) Recv() (*audio_api_v1.CreateEnrollmentResponse, error) {
	m := new(audio_api_v1.CreateEnrollmentResponse)
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

// Mock Authenticate Client
type mockBiometricsAuthenticateClient struct{}

func (c *mockBiometricsAuthenticateClient) Send(m *audio_api_v1.AuthenticateRequest) error {
	return c.SendMsg(m)
}

func (c *mockBiometricsAuthenticateClient) Recv() (*audio_api_v1.AuthenticateResponse, error) {
	m := new(audio_api_v1.AuthenticateResponse)
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

// Mock Events Client
type mockEventsClient struct{}

func (c *mockEventsClient) ValidateEvent(ctx context.Context, opts ...grpc.CallOption) (audio_api_v1.AudioEvents_ValidateEventClient, error) {
	validateEventClient := &mockValidateEventClient{}
	return validateEventClient, nil
}

func (c *mockEventsClient) CreateEnrolledEvent(ctx context.Context, opts ...grpc.CallOption) (audio_api_v1.AudioEvents_CreateEnrolledEventClient, error) {
	createEnrolledEventClient := &mockCreateEnrolledEventClient{}
	return createEnrolledEventClient, nil
}

func (c *mockEventsClient) ValidateEnrolledEvent(ctx context.Context, opts ...grpc.CallOption) (audio_api_v1.AudioEvents_ValidateEnrolledEventClient, error) {
	validateEnrolledEventClient := &mockValidateEnrolledEventClient{}
	return validateEnrolledEventClient, nil
}

// Mock Validate Event Client
type mockValidateEventClient struct{}

func (c *mockValidateEventClient) Send(m *audio_api_v1.ValidateEventRequest) error {
	return c.SendMsg(m)
}

func (c *mockValidateEventClient) Recv() (*audio_api_v1.ValidateEventResponse, error) {
	m := new(audio_api_v1.ValidateEventResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockValidateEventClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockValidateEventClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockValidateEventClient) CloseSend() error {
	return nil
}

func (stream *mockValidateEventClient) Context() context.Context {
	return context.Background()
}

func (stream *mockValidateEventClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockValidateEventClient) RecvMsg(m interface{}) error {
	return nil
}

// Mock Create Enrolled Event Client
type mockCreateEnrolledEventClient struct{}

func (c *mockCreateEnrolledEventClient) Send(m *audio_api_v1.CreateEnrolledEventRequest) error {
	return c.SendMsg(m)
}

func (c *mockCreateEnrolledEventClient) Recv() (*audio_api_v1.CreateEnrollmentResponse, error) {
	m := new(audio_api_v1.CreateEnrollmentResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockCreateEnrolledEventClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockCreateEnrolledEventClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockCreateEnrolledEventClient) CloseSend() error {
	return nil
}

func (stream *mockCreateEnrolledEventClient) Context() context.Context {
	return context.Background()
}

func (stream *mockCreateEnrolledEventClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockCreateEnrolledEventClient) RecvMsg(m interface{}) error {
	return nil
}

// Mock Validate Enrolled Event Client
type mockValidateEnrolledEventClient struct{}

func (c *mockValidateEnrolledEventClient) Send(m *audio_api_v1.ValidateEnrolledEventRequest) error {
	return c.SendMsg(m)
}

func (c *mockValidateEnrolledEventClient) Recv() (*audio_api_v1.ValidateEnrolledEventResponse, error) {
	m := new(audio_api_v1.ValidateEnrolledEventResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockValidateEnrolledEventClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockValidateEnrolledEventClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockValidateEnrolledEventClient) CloseSend() error {
	return nil
}

func (stream *mockValidateEnrolledEventClient) Context() context.Context {
	return context.Background()
}

func (stream *mockValidateEnrolledEventClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockValidateEnrolledEventClient) RecvMsg(m interface{}) error {
	return nil
}

// Mock Audio Models Client
type mockAudioModelsClient struct{}

func (c *mockAudioModelsClient) GetModels(ctx context.Context, in *audio_api_v1.GetModelsRequest, opts ...grpc.CallOption) (*audio_api_v1.GetModelsResponse, error) {
	model := &audio_api_v1.AudioModel{
		Name:                "test-model",
		IsEnrollable:        false,
		ModelType:           common.ModelType_VOICE_BIOMETRIC_WAKEWORD,
		FixedPhrase:         "fixed phrase",
		SampleRate:          16000,
		Versions:            []string{"version0"},
		Technology:          common.TechnologyType_NOT_SET,
		IsLivenessSupported: false,
	}
	return &audio_api_v1.GetModelsResponse{Models: []*audio_api_v1.AudioModel{model}}, nil
}

// Mock Audio Transcription Client
type mockAudioTranscriptionClient struct{}

func (c *mockAudioTranscriptionClient) Transcribe(ctx context.Context, opts ...grpc.CallOption) (audio_api_v1.AudioTranscriptions_TranscribeClient, error) {
	transcribeClient := &mockTranscribeClient{}
	return transcribeClient, nil
}

// Mock Transcribe Client
type mockTranscribeClient struct{}

func (c *mockTranscribeClient) Send(m *audio_api_v1.TranscribeRequest) error {
	return c.SendMsg(m)
}

func (c *mockTranscribeClient) Recv() (*audio_api_v1.TranscribeResponse, error) {
	m := new(audio_api_v1.TranscribeResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockTranscribeClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockTranscribeClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockTranscribeClient) CloseSend() error {
	return nil
}

func (stream *mockTranscribeClient) Context() context.Context {
	return context.Background()
}

func (stream *mockTranscribeClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockTranscribeClient) RecvMsg(m interface{}) error {
	return nil
}

// Mock Audio Synthesis Client
type mockAudioSynthesisClient struct{}

func (c *mockAudioSynthesisClient) SynthesizeSpeech(ctx context.Context, in *audio_api_v1.SynthesizeSpeechRequest, opts ...grpc.CallOption) (
	audio_api_v1.AudioSynthesis_SynthesizeSpeechClient,
	error,
) {
	synthesizeSpeechClient := &mockSynthesizeSpeechClient{}
	return synthesizeSpeechClient, nil
}

// Mock Synthesize Speech Client
type mockSynthesizeSpeechClient struct{}

func (c *mockSynthesizeSpeechClient) Recv() (*audio_api_v1.SynthesizeSpeechResponse, error) {
	m := new(audio_api_v1.SynthesizeSpeechResponse)
	if err := c.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (stream *mockSynthesizeSpeechClient) Header() (metadata.MD, error) {
	return metadata.MD{}, nil
}

func (stream *mockSynthesizeSpeechClient) Trailer() metadata.MD {
	return metadata.MD{}
}

func (stream *mockSynthesizeSpeechClient) CloseSend() error {
	return nil
}

func (stream *mockSynthesizeSpeechClient) Context() context.Context {
	return context.Background()
}

func (stream *mockSynthesizeSpeechClient) SendMsg(m interface{}) error {
	return nil
}

func (stream *mockSynthesizeSpeechClient) RecvMsg(m interface{}) error {
	return nil
}
