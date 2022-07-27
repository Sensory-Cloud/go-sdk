# Go SDK SDK

Sensory Cloud Go SDK documentation.

## General Information

Before getting started, you must spin up a Sensory Cloud inference server or have Sensory spin one up for you. You must also have the following pieces of information:
* Your inference server URL
* Your Sensory Tenant ID (UUID)
* Your configured secret key used to register OAuth clients

## Checking Server Health

It's important to check the health of your Sensory Inference server. You can do so via the following:

```go
// Tenant ID granted by Sensory Inc.
sensoryTenantId := "f6580f3b-dcaf-465b-867e-59fbbb0ab3fc"
deviceId := "a-hardware-identifier-unique-to-your-device"

// Configuration specific to your tenant
config := config.ClientConfig{
    FullyQualifiedDomainName: "your-inference-server.com",
    TenantId:                 sensoryTenantId,
    DeviceId:                 deviceId,
}

onClose, err := config.Connect()
if err != nil {
    panic(err)
}
defer onClose()


healthService, err := health_service.NewHealthService(&config)
if err != nil {
    panic(err)
}

// Create context for the grpc request
ctx := context.Background()
// Optionally, set a timeout on the context
ctx, cancel := context.WithTimeout(ctx, time.Minute)
defer cancel()

serverHealth, err := healthService.GetHealth(ctx)
if err != nil {
    panic(err)
}
```

## Secure Credential Store

ISecureCredential is an interface that store and serves your OAuth credentials (clientId and clientSecret).
ISecureCredential must be implemented by you and the credentials should be persisted in a secure manner, such as in an encrypted database.
OAuth credentials should be generated one time per unique machine.

A crude example of ISecureCredential is as follows:

```go
type SecureCredentialStore struct {
	ClientId string
	ClientSecret string
}

func (s SecureCredentialStore) SaveCredentials(clientId string, secret string) {
    s.ClientId = clientId
    s.ClientSecret = secret
}

func (s SecureCredentialStore) GetClientId() string {
	return s.ClientId
}

func (s SecureCredentialStore) GetClientSecret() string {
	return s.ClientSecret
}
```

## Registering OAuth Credentials

OAuth credentials should be registered once per unique machine. Registration is very simple, and provided as part of the SDK.

The below example shows how to create an OAuthService and register a client for the first time.

```go
// Tenant ID granted by Sensory Inc.
sensoryTenantId := "f6580f3b-dcaf-465b-867e-59fbbb0ab3fc"
deviceId := "a-hardware-identifier-unique-to-your-device"
deviceName := "a-user-friendly-name-for-this-device"

// Configuration specific to your tenant
// Configurations may also be loaded from a file name
//  - see `pkg/file_parser/testdata` for example configuration files
sdkInitConfig := config.SDKInitConfig{
    FullyQualifiedDomainName: "your-inference-server.com",
    IsSecure: true,
    TenantID: sensoryTenantId,
    EnrollmentType: config.SharedSecret,
    Credential: "password",
    DeviceID: deviceId,
    DeviceName: deviceName,
}

// Create a secure credential store
credentialStore := SecureCredentialStore{}

// Create context for the grpc request
ctx := context.Background()
// Optionally, set a timeout on the context
ctx, cancel := context.WithTimeout(ctx, time.Minute)
defer cancel()

// Initialize the SDK
sdkInitializer := initializer.NewInitializer(credentialStore, nil)
clientConfig, response, err := sdkInitializer.Initialize(context, sdkInitConfig)
if err != nil {
    panic(err)
}

// Connect to the returned client config
// The clientConfig should be saved and used when initializing other Sensory Cloud services
onClose, err := clientConfig.Connect()
if err != nil {
    panic(err)
}
defer onClose()


// OAuth registration can take one of two paths, the unsecure path that uses a shared secret between this device and your instance of Sensory Cloud
// or asymmetric public / private keypair registration.

// Path 1 (used in the above example) --------
// Insecure authorization credential as configured on your instance of Sensory Cloud
sdkInitConfig := config.SDKInitConfig{
    FullyQualifiedDomainName: "your-inference-server.com",
    IsSecure: true,
    TenantID: sensoryTenantId,
    EnrollmentType: config.SharedSecret,
    Credential: "password",
    DeviceID: deviceId,
    DeviceName: deviceName,
}

// Path 2 --------
// Secure Public / private keypair registration using Portable.BouncyCastle and ScottBrady.IdentityModel

sdkInitConfig := config.SDKInitConfig{
    FullyQualifiedDomainName: "your-inference-server.com",
    IsSecure: true,
    TenantID: sensoryTenantId,
    EnrollmentType: config.jwt,
    Credential: "your-signing-private-key",
    DeviceID: deviceId,
    DeviceName: deviceName,
}

// Keypair generation happens out of band and long before a registration occurs. The below example shows how to generate and sign an enrollment JWT as a comprehensive example.
// For this path, a JWT signer must be passed into the SDK initializer that can create a signed JWT

// Public key is persisted in Sensory Cloud via the Management interface
publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
if err != nil {
    panic(err)
}

claims := jwt.Claims{}
token, err := claims.EdDSASign(privateKey)
if err != nil {
    panic(err)
}
```

## Renewing Credentials

If you are evaluating Sensory Cloud as part of a Proof of Concept then likely the token granted to you will expire.
Should the secret expire, you can request a new credentials from from Sensory. You will need to renew your credential via the OAuth service.
You can always renew your secret before it has expired to ensure you have 0 downtime.

```go
oauthService.RenewDeviceCredential(newToken);
```


## Creating a TokenManager

The TokenManger class handles requesting OAuth tokens when necessary.

```go
clientConfig := config.ClientConfig{} // Use clientConfig returned by the initialize call

onClose, err := clientConfig.Connect()
if err != nil {
    panic(err)
}
defer onClose()

credentialStore := SecureCredentialStore{}
oauthService, err := NewOauthService(clientConfig, credentialStore);
if err != nil {
    panic(err)
}
tokenManager := NewTokenManager(oauthService);

// Obtain OAuth token from Sensory Cloud
string token = tokenManager.GetToken();

return tokenManager;

```

## Creating an AudioService

AudioService provides methods to stream audio to Sensory Cloud. It is recommended to only have 1 instance of AudioService
instantiated per Config. In most circumstances you will only ever have one Config, unless your app communicates with
multiple Sensory Cloud servers.

```go
clientConfig := config.ClientConfig{} // Use clientConfig returned by the initialize call

onClose, err := clientConfig.Connect()
if err != nil {
    panic(err)
}
defer onClose()

credentialStore := SecureCredentialStore{}
oauthService, err := NewOauthService(clientConfig, credentialStore);
if err != nil {
    panic(err)
}
tokenManager := NewTokenManager(oauthService);
audioService := NewAudioService(clientConfig, tokenManager)

```

### Obtaining Audio Models

Certain audio models are available to your application depending on the models that are configured for your instance of Sensory Cloud.
In order to determine which audio models are accessible to you, you can execute the below code.

```go
audioService := audio_service.NewAudioService(config, tokenManager)

// Create context for the grpc request
ctx := context.Background()
// Optionally, set a timeout on the context
ctx, cancel := context.WithTimeout(ctx, time.Minute)
defer cancel()

audioModels := audioService.GetModels(ctx);
```

Audio models contain the following properties:
* Name - the unique name tied to this model. Used when calling any other audio function.
* IsEnrollable - indicates if the model can be enrolled into. Models that are enrollable can be used in the CreateEnrollment function.
* ModelType - indicates the class of model and it's general function.
* FixedPhrase - for speech-based models only. Indicates if a specific phrase must be said.
* Samplerate - indicates the audio samplerate required by this model. Generally, the number will be 16000.
* IsLivenessSupported - indicates if this model supports liveness for enrollment and authentication. Liveness provides an added layer of security by requiring a users to speak random digits.

### Enrolling with Audio

In order to enroll with audio, you must first ensure you have an enrollable model enabled for your Sensory Cloud instance. This can be obtained via the GetModels() request.
Enrolling with audio uses a bi-directional streaming pattern to allow immediate feedback to the user during enrollment. It is important to save the enrollmentId
in order to perform authentication against it in the future.

```go
audioService := NewAudioService(config, tokenManager)
userId := "user1@test.com"
modelName := "wakeword-16kHz-open_sesame.ubm"
isLivenessEnabled := false

audioConfig := &audio_api_v1.CreateEnrollmentConfig{
    Audio: &audio_api_v1.AudioConfig{
        Encoding:          audio_api_v1.AudioConfig_LINEAR16,
        AudioChannelCount: 1,
        SampleRateHertz:   16000,
        LanguageCode:      "en-US",
    },
    Description:       "my enrollment",
    UserId:            userId,
    ModelName:         modelName,
    IsLivenessEnabled: isLivenessEnabled,
}

// Create context for the grpc request
ctx := context.Background()
// Optionally, set a timeout on the context
ctx, cancel := context.WithTimeout(ctx, time.Minute)
defer cancel()

stream, err := audioSerice.StreamEnrollment(ctx, audioConfig)
defer stream.CloseSend()
test_util.AssertOk(t, err)

// Create channels to handle communication between goroutines
errorChan := make(chan error)
responseChan := make(chan *audio_api_v1.CreateEnrollmentResponse, 10)

// Channel to pass audio into (microphone or otherwise)
audioChan := make(chan []byte, 3)

// Listen for responses on the gRPC pipe on a separate goroutine

go func() {
    for {
        response, err := stream.Recv()
        if err != nil {
            errorChan <- err
            return
        }

        responseChan <- response
    }
}()

for {
    select {
    // Watch for errors
    case err := <-errorChan:
        panic(err)

    // Or timeout after 1 minute
    case <-time.NewTimer(time.Minute).C:
        panic("read timeout!")

    // Read responses on the main goroutine
    case result := <-responseChan:
        fmt.Printf("Response from server: %v", result)
        if result.EnrollmentId != "nil" {
            fmt.Printf("Enrollment Complete with ID %s", result.EnrollmentId)
            return
        }

    // Send audio bytes up the pipe
    case audio := <-audioChan:
        err := stream.Send(&audio_api_v1.CreateEnrollmentRequest{
            StreamingRequest: &audio_api_v1.CreateEnrollmentRequest_AudioContent{
                AudioContent: audio,
            },
        })
        if err != nil {
            panic(err)
        }
    }
}
```

### Authenticating with Audio

Authenticating with audio is similar to enrollment, except now you have an enrollmentId to pass into the function.

```go
// TODO
```

### Audio events

Audio events are used to recognize specific words, phrases, or sounds.

The below example waits for a single event to be recognized and ends the stream.

```go
// TODO
```

The below example outlines how to stream a wav file in chunks to Sensory Cloud and process all events before closing the stream.

```go
// TODO
```

## CreateEnrolledEvent

You can enroll your own event into the Sensory cloud system. The process is similar to biometric enrollment where you must
play a sound or speak a particular phrase 4 or more times. This is usefull for recognizing sounds that are not offered by Sensory Cloud.

```go
// TODO

```

### ValidateEnrolledEvent

Once you've created an enroled event, you can listen for that event (or groups of events) by calling
the ValidateEnrolledEvent function.

```go
// TODO
```

### Transcription

Transcription is used to convert audio into text.

```go
audioSerice, err := audio_service.NewAudioService(config, tokenManager)
userId := "user1@test.com"
modelName := "wakeword-16kHz-open_sesame.ubm"

audioConfig := &audio_api_v1.TranscribeConfig{
    Audio: &audio_api_v1.AudioConfig{
        Encoding:          audio_api_v1.AudioConfig_LINEAR16,
        AudioChannelCount: 1,
        SampleRateHertz:   16000,
        LanguageCode:      "en-US",
    },
    UserId:    userId,
    ModelName: modelName,
}

aggregator := audio_service.FullTranscriptAggregator{}

// Create context for the grpc request
ctx := context.Background()
// Optionally, set a timeout on the context
ctx, cancel := context.WithTimeout(ctx, time.Minute)
defer cancel()

stream, err := audioSerice.StreamTranscription(ctx, audioConfig)
defer stream.CloseSend()
test_util.AssertOk(t, err)

// Create channels to handle communication between goroutines
errorChan := make(chan error)
responseChan := make(chan *audio_api_v1.TranscribeResponse, 10)

// Channel to pass audio into (microphone or otherwise)
audioChan := make(chan []byte, 3)

// Listen for responses on the gRPC pipe on a separate goroutine
go func() {
    for {
        response, err := stream.Recv()
        if err != nil {
            errorChan <- err
            return
        }

        responseChan <- response
    }
}()

for {
    select {
    // Watch for errors
    case err := <-errorChan:
        panic(err)

    // Or timeout after 1 minute
    case <-time.NewTimer(time.Minute).C:
        panic("read timeout!")

    // Read responses on the main goroutine
    case result := <-responseChan:
        fmt.Printf("Response from server: %v", result)

        // Process the response using the aggregator
        aggregator.ProcessResponse(response.GetWordList())

    // Send audio bytes up the pipe
    case audio := <-audioChan:
        err := stream.Send(&audio_api_v1.TranscribeRequest{
            StreamingRequest: &audio_api_v1.TranscribeRequest_AudioContent{
                AudioContent: audio,
            },
        })
        if err != nil {
            panic(err)
        }
    }
}

// aggregator.GetCurrentTranscript() is the entire transcript as a string
// aggregator.GetCurrentWordList() is every word stored in a slice with metadata such as start, end, and confidence of each word
```

## Creating a VideoService

VideoService provides methods to stream images to Sensory Cloud. It is recommended to only have 1 instance of VideoService
instantiated per Config. In most circumstances you will only ever have one Config, unless your app communicates with
multiple Sensory Cloud servers.

```go
// TODO
```

### Obtaining Video Models

Certain video models are available to your application depending on the models that are configured for your instance of Sensory Cloud.
In order to determine which video models are accessible to you, you can execute the below code.

```go
// TODO
```

Audio models contain the following properties:
* Name - the unique name tied to this model. Used when calling any other video function.
* IsEnrollable - indicates if the model can be enrolled into. Models that are enrollable can be used in the CreateEnrollment function.
* ModelType - indicates the class of model and it's general function.
* FixedObject - for recognition-based models only. Indicates if this model is built to recognize a specific object.
* IsLivenessSupported - indicates if this model supports liveness for enrollment and authentication. Liveness provides an added layer of security.

### Enrolling with Video

In order to enroll with video, you must first ensure you have an enrollable model enabled for your Sensory Cloud instance. This can be obtained via the GetModels() request.
Enrolling with video uses a call and response streaming pattern to allow immediate feedback to the user during enrollment. It is important to save the enrollmentId
in order to perform authentication against it in the future.

```go
// TODO
```

### Authenticating with Video

Authenticating with video is similar to enrollment, except now you have an enrollmentId to pass into the function.

```go
// TODO
```

## Creating a ManagementService

The ManagementService is used to manage typical CRUD operations with Sensory Cloud, such as deleting enrollments or creating enrollment groups.
For more information on the specific functions of the ManagementService, please refer to the ManagementService.cs file located in the Services folder.

```go
clientConfig := config.ClientConfig{} // Use clientConfig returned by the initialize call

onClose, err := clientConfig.Connect()
if err != nil {
    panic(err)
}
defer onClose()

credentialStore := SecureCredentialStore{}
oauthService, err := NewOauthService(clientConfig, credentialStore);
if err != nil {
    panic(err)
}
tokenManager := NewTokenManager(oauthService);

managementService, err := NewManagementService(clientConfig, tokenManager);
if err != nil {
    panic(err)
}
```
