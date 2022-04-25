package video_service

import (
	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	"github.com/Sensory-Cloud/go-sdk/pkg/config"
	video_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/video"
	token_manager "github.com/Sensory-Cloud/go-sdk/pkg/token"
)

type VideoService struct {
	config                 *config.ClientConfig
	tokenManager           token_manager.IAuthorizationMetadataService
	videoBiometricsClient  video_api_v1.VideoBiometricsClient
	videoModelsClient      video_api_v1.VideoModelsClient
	videoRecognitionClient video_api_v1.VideoRecognitionClient
}

// NewVideoService creates a service to handle image requests to Sensory Cloud
func NewVideoService(config *config.ClientConfig, tokenManager token_manager.IAuthorizationMetadataService) (*VideoService, error) {
	client, err := config.GetClient()
	if err != nil {
		return nil, err
	}

	videoBiometricsClient := video_api_v1.NewVideoBiometricsClient(client)
	videoModelsClient := video_api_v1.NewVideoModelsClient(client)
	videoRecognitionClient := video_api_v1.NewVideoRecognitionClient(client)
	return &VideoService{config, tokenManager, videoBiometricsClient, videoModelsClient, videoRecognitionClient}, nil
}

//  Fetch all the vision models supported by your instance of Sensory Cloud.
func (s *VideoService) GetModels() (*video_api_v1.GetModelsResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.videoModelsClient.GetModels(ctx, &video_api_v1.GetModelsRequest{})
}

// Stream images to Sensory Cloud as a means for user enrollment.
// Only biometric-typed models are supported by the method.
func (s *VideoService) StreamEnrollment(config *video_api_v1.CreateEnrollmentConfig) (video_api_v1.VideoBiometrics_CreateEnrollmentClient, error) {

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

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
func (s *VideoService) StreamAuthentication(config *video_api_v1.AuthenticateConfig) (video_api_v1.VideoBiometrics_AuthenticateClient, error) {

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

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
func (s *VideoService) StreamLivenessRecognition(config *video_api_v1.ValidateRecognitionConfig) (video_api_v1.VideoRecognition_ValidateLivenessClient, error) {

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

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
