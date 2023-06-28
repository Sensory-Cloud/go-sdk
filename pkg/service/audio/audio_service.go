package audio_service

import (
	"context"

	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	audio_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/audio"
	token_manager "github.com/Sensory-Cloud/go-sdk/pkg/token"
)

type AudioService struct {
	config                    config.IClientConfig
	tokenManager              token_manager.IAuthorizationMetadataService
	audioBiometricsClient     audio_api_v1.AudioBiometricsClient
	audioEventsClient         audio_api_v1.AudioEventsClient
	audioModelsClient         audio_api_v1.AudioModelsClient
	audioTranscriptionsClient audio_api_v1.AudioTranscriptionsClient
	audioSynthesisClient      audio_api_v1.AudioSynthesisClient
}

// NewAudioService creates a service to handle audio requests to Sensory Cloud
func NewAudioService(config config.IClientConfig, tokenManager token_manager.IAuthorizationMetadataService) (*AudioService, error) {
	client, err := config.GetClient()
	if err != nil {
		return nil, err
	}

	audioBiometricsClient := audio_api_v1.NewAudioBiometricsClient(client)
	audioEventsClient := audio_api_v1.NewAudioEventsClient(client)
	audioModelsClient := audio_api_v1.NewAudioModelsClient(client)
	audioTranscriptionsClient := audio_api_v1.NewAudioTranscriptionsClient(client)
	audioSynthesisClient := audio_api_v1.NewAudioSynthesisClient(client)
	return &AudioService{config, tokenManager, audioBiometricsClient, audioEventsClient, audioModelsClient, audioTranscriptionsClient, audioSynthesisClient}, nil
}

// Fetch all the audio models supported by your instance of Sensory Cloud.
func (s *AudioService) GetModels(ctx context.Context) (*audio_api_v1.GetModelsResponse, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.audioModelsClient.GetModels(ctx, &audio_api_v1.GetModelsRequest{})
}

// Stream audio to Sensory Cloud as a means for user enrollment.
// Only biometric-typed models are supported by the method.
func (s *AudioService) StreamEnrollment(ctx context.Context, config *audio_api_v1.CreateEnrollmentConfig) (audio_api_v1.AudioBiometrics_CreateEnrollmentClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	enrollmentStream, err := s.audioBiometricsClient.CreateEnrollment(ctx)
	if err != nil {
		return nil, err
	}

	err = enrollmentStream.Send(&audio_api_v1.CreateEnrollmentRequest{
		StreamingRequest: &audio_api_v1.CreateEnrollmentRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return enrollmentStream, nil
}

// Authenticate against an existing audio enrollment in Sensory Cloud.
// Only biometric-typed models are supported by the method.
func (s *AudioService) StreamAuthentication(ctx context.Context, config *audio_api_v1.AuthenticateConfig) (audio_api_v1.AudioBiometrics_AuthenticateClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	authenticationStream, err := s.audioBiometricsClient.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	err = authenticationStream.Send(&audio_api_v1.AuthenticateRequest{
		StreamingRequest: &audio_api_v1.AuthenticateRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return authenticationStream, nil
}

// Stream audio to Sensory Cloud as a means for audio enrollment.
// This method is similar to StreamEnrollment, but does not have the same
// time limits or model type restrictions.
// Biometric model types are not supported by this function.
// This endpoint cannot be used to establish device trust.
func (s *AudioService) StreamCreateEnrolledEvent(ctx context.Context, config *audio_api_v1.CreateEnrollmentEventConfig) (audio_api_v1.AudioEvents_CreateEnrolledEventClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	enrollmentStream, err := s.audioEventsClient.CreateEnrolledEvent(ctx)
	if err != nil {
		return nil, err
	}

	err = enrollmentStream.Send(&audio_api_v1.CreateEnrolledEventRequest{
		StreamingRequest: &audio_api_v1.CreateEnrolledEventRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return enrollmentStream, nil
}

// Stream audio to Sensory Cloud in order to recognize a specific phrase or sound
func (s *AudioService) StreamEvent(ctx context.Context, config *audio_api_v1.ValidateEventConfig) (audio_api_v1.AudioEvents_ValidateEventClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	eventStream, err := s.audioEventsClient.ValidateEvent(ctx)
	if err != nil {
		return nil, err
	}

	err = eventStream.Send(&audio_api_v1.ValidateEventRequest{
		StreamingRequest: &audio_api_v1.ValidateEventRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return eventStream, nil
}

// Validate an existing event enrollment in Sensory Cloud.
// This method is similar to Authenticate, but does not have the same
// time limits or model type restrictions. Additionally, the server will
// never close the stream, and thus a client may validate an enrolled sound
// as many times as they'd like.
// Any model types are supported by this function.
// This endpoint cannot be used to establish device trust.
func (s *AudioService) StreamValidateEnrolledEvent(ctx context.Context, config *audio_api_v1.ValidateEnrolledEventConfig) (audio_api_v1.AudioEvents_ValidateEnrolledEventClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	validateEnrolledEventStream, err := s.audioEventsClient.ValidateEnrolledEvent(ctx)
	if err != nil {
		return nil, err
	}

	err = validateEnrolledEventStream.Send(&audio_api_v1.ValidateEnrolledEventRequest{
		StreamingRequest: &audio_api_v1.ValidateEnrolledEventRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return validateEnrolledEventStream, nil
}

// Stream audio to Sensory Cloud in order to transcribe spoken words
func (s *AudioService) StreamTranscription(ctx context.Context, config *audio_api_v1.TranscribeConfig) (audio_api_v1.AudioTranscriptions_TranscribeClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	transcriptionsStream, err := s.audioTranscriptionsClient.Transcribe(ctx)
	if err != nil {
		return nil, err
	}

	err = transcriptionsStream.Send(&audio_api_v1.TranscribeRequest{
		StreamingRequest: &audio_api_v1.TranscribeRequest_Config{
			Config: config,
		},
	})
	if err != nil {
		return nil, err
	}

	return transcriptionsStream, nil
}

// Streams synthesized speech back from Sensory Cloud
// Concatenating all of the `audioContent` of the responses from the returned client will result in a complete WAV file of the resultant audio
func (s *AudioService) SynthesizeSpeech(ctx context.Context, request *audio_api_v1.SynthesizeSpeechRequest) (audio_api_v1.AudioSynthesis_SynthesizeSpeechClient, error) {
	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.audioSynthesisClient.SynthesizeSpeech(ctx, request)
}

// Set client config
func (s *AudioService) SetConfig(config config.IClientConfig) {
	s.config = config
}

// Set token manager
func (s *AudioService) SetTokenManager(tokenManager token_manager.IAuthorizationMetadataService) {
	s.tokenManager = tokenManager
}

// Set audio biometrics client
func (s *AudioService) SetAudioBiometricsClient(client audio_api_v1.AudioBiometricsClient) {
	s.audioBiometricsClient = client
}

// Set audio events client
func (s *AudioService) SetAudioEventsClient(client audio_api_v1.AudioEventsClient) {
	s.audioEventsClient = client
}

// Set audio models client
func (s *AudioService) SetAudioModelsClient(client audio_api_v1.AudioModelsClient) {
	s.audioModelsClient = client
}

// Set audio transcriptions client
func (s *AudioService) SetAudioTranscriptionsClient(client audio_api_v1.AudioTranscriptionsClient) {
	s.audioTranscriptionsClient = client
}

// Set audio synthesis client
func (s *AudioService) SetAudioSynthesisClient(client audio_api_v1.AudioSynthesisClient) {
	s.audioSynthesisClient = client
}
