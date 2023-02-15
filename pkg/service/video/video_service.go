package video_service

import (
	"context"

	"github.com/Sensory-Cloud/go-sdk/pkg/config"
	video_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/video"
	token_manager "github.com/Sensory-Cloud/go-sdk/pkg/token"
)

type VideoService struct {
	config                 config.IClientConfig
	tokenManager           token_manager.IAuthorizationMetadataService
	videoBiometricsClient  video_api_v1.VideoBiometricsClient
	videoModelsClient      video_api_v1.VideoModelsClient
	videoRecognitionClient video_api_v1.VideoRecognitionClient
}

// NewVideoService creates a service to handle image requests to Sensory Cloud
func NewVideoService(config config.IClientConfig, tokenManager token_manager.IAuthorizationMetadataService) (*VideoService, error) {
	client, err := config.GetClient()
	if err != nil {
		return nil, err
	}

	videoBiometricsClient := video_api_v1.NewVideoBiometricsClient(client)
	videoModelsClient := video_api_v1.NewVideoModelsClient(client)
	videoRecognitionClient := video_api_v1.NewVideoRecognitionClient(client)
	return &VideoService{config, tokenManager, videoBiometricsClient, videoModelsClient, videoRecognitionClient}, nil
}

// Fetch all the vision models supported by your instance of Sensory Cloud.
func (s *VideoService) GetModels(ctx context.Context) (*video_api_v1.GetModelsResponse, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.videoModelsClient.GetModels(ctx, &video_api_v1.GetModelsRequest{})
}

// Stream images to Sensory Cloud as a means for user enrollment.
// Only biometric-typed models are supported by the method.
func (s *VideoService) StreamEnrollment(ctx context.Context, config *video_api_v1.CreateEnrollmentConfig) (video_api_v1.VideoBiometrics_CreateEnrollmentClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	enrollmentStream, err := s.videoBiometricsClient.CreateEnrollment(ctx)
	if err != nil {
		return nil, err
	}

	err = enrollmentStream.Send(&video_api_v1.CreateEnrollmentRequest{
		StreamingRequest: &video_api_v1.CreateEnrollmentRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return enrollmentStream, nil
}

// Authenticate against an existing enrollment in Sensory Cloud.
// Only biometric-typed models are supported by the method.
func (s *VideoService) StreamAuthentication(ctx context.Context, config *video_api_v1.AuthenticateConfig) (video_api_v1.VideoBiometrics_AuthenticateClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	authenticationStream, err := s.videoBiometricsClient.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	err = authenticationStream.Send(&video_api_v1.AuthenticateRequest{
		StreamingRequest: &video_api_v1.AuthenticateRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return authenticationStream, nil
}

// Stream images to Sensory Cloud in order to recognize "liveness" of a particular image.
func (s *VideoService) StreamLivenessRecognition(ctx context.Context, config *video_api_v1.ValidateRecognitionConfig) (video_api_v1.VideoRecognition_ValidateLivenessClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	authenticationStream, err := s.videoRecognitionClient.ValidateLiveness(ctx)
	if err != nil {
		return nil, err
	}

	err = authenticationStream.Send(&video_api_v1.ValidateRecognitionRequest{
		StreamingRequest: &video_api_v1.ValidateRecognitionRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return authenticationStream, nil
}

// Set client config
func (s *VideoService) SetClientConfig(config config.IClientConfig) {
	s.config = config
}

// Set token manager
func (s *VideoService) SetTokenManager(tokenManager token_manager.IAuthorizationMetadataService) {
	s.tokenManager = tokenManager
}

// Set video biometrics client
func (s *VideoService) SetVideoBiometricsClient(client video_api_v1.VideoBiometricsClient) {
	s.videoBiometricsClient = client
}

// Set video models cient
func (s *VideoService) SetVideoModelsClient(client video_api_v1.VideoModelsClient) {
	s.videoModelsClient = client
}

// Set video recognition client
func (s *VideoService) SetVideoRecognitionClient(client video_api_v1.VideoRecognitionClient) {
	s.videoRecognitionClient = client
}
