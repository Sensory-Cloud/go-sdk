package audio_service_test

import (
	"testing"
	// audio_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/audio"
	// audio_service "github.com/Sensory-Cloud/go-sdk/pkg/service/audio"
	// test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestNewAudioService(t *testing.T) {
	t.FailNow()

	// audioSerice, err := audio_service.NewAudioService(config, tokenManager)
	// userId := "user1@test.com"
	// modelName := "wakeword-16kHz-open_sesame.ubm"

	// audioConfig := &audio_api_v1.TranscribeConfig{
	// 	Audio: &audio_api_v1.AudioConfig{
	// 		Encoding:          audio_api_v1.AudioConfig_LINEAR16,
	// 		AudioChannelCount: 1,
	// 		SampleRateHertz:   16000,
	// 		LanguageCode:      "en-US",
	// 	},
	// 	UserId:    userId,
	// 	ModelName: modelName,
	// }

	// // Create context for the grpc request
	// ctx := context.Background()
	// // Optionally, set a timeout on the context
	// ctx, cancel := context.WithTimeout(ctx, time.Minute)
	// defer cancel()

	// stream, err := audioSerice.StreamTranscription(ctx, audioConfig)
	// defer stream.CloseSend()
	// test_util.AssertOk(t, err)

	// // Create channels to handle communication between goroutines
	// errorChan := make(chan error)
	// responseChan := make(chan *audio_api_v1.TranscribeResponse, 10)

	// // Channel to pass audio into (microphone or otherwise)
	// audioChan := make(chan []byte, 3)

	// // Listen for responses on the gRPC pipe on a separate goroutine
	// go func() {
	// 	response, err := stream.Recv()
	// 	if err != nil {
	// 		errorChan <- err
	// 	}

	// 	responseChan <- response
	// }()

	// for {
	// 	select {
	// 	// Watch for errors
	// 	case err := <-errorChan:
	// 		panic(err)

	// 	// Or timeout after 1 minute
	// 	case <-time.NewTimer(time.Minute).C:
	// 		panic("read timeout!")

	// 	// Read responses on the main goroutine
	// 	case result := <-responseChan:
	// 		fmt.Printf("Response from server: %v", result)

	// 	// Send audio bytes up the pipe
	// 	case audio := <-audioChan:
	// 		err := stream.Send(&audio_api_v1.TranscribeRequest{
	// 			StreamingRequest: &audio_api_v1.TranscribeRequest_AudioContent{
	// 				AudioContent: audio,
	// 			},
	// 		})
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
}

func TestStreamEnrollment(t *testing.T) {
	t.FailNow()
}

func TestStreamAuthentication(t *testing.T) {
	t.FailNow()
}

func TestStreamCreateEnrolledEvent(t *testing.T) {
	t.FailNow()
}

func TestStreamEvent(t *testing.T) {
	t.FailNow()
}

func TestStreamValidateEnrolledEvent(t *testing.T) {
	t.FailNow()
}

func TestStreamTranscription(t *testing.T) {
	t.FailNow()
}

func TestSynthesizeSpeech(t *testing.T) {
	t.FailNow()
}
