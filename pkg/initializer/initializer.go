package initializer

import (
	"context"

	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	file_parser "github.com/Sensory-Cloud/go-sdk/pkg/file_parser"
	management_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/management"
	oauth_service "github.com/Sensory-Cloud/go-sdk/pkg/service/oauth"
)

var (
	// Initializer function to create a new oauth service. Should only be changed for unit testing
	OauthServiceInit func(config.IClientConfig, oauth_service.ISecureCredentialStore) (oauth_service.IOauthService, error) = oauth_service.NewOauthService
)

// Creates a signed JWT for device enrollment
type IJWTSigner interface {
	// Creates a signed JWT used for device enrollment
	// This is only used if the enrollment type is `jwt`
	// Parameters:
	//  - enrollmentKey: private signing key. This will be a hex string of the raw private key
	//  - deviceName: device name to use in JWT
	//  - tenantID: tenant ID to use in JWT
	//  - clientID: client ID to use in JWT
	SignJWT(enrollmentKey string, deviceName string, tenantID string, clientID string) (string, error)
}

// Initializes the Sensory Cloud SDK and automatically registers the device if not already done so
type Initializer struct {
	secureCredentialStore oauth_service.ISecureCredentialStore
	jwtSigner             IJWTSigner
}

// Creates a new initializer object
// Parameters:
//   - credentialStore: Secure credential store for oauth credentials
//   - jwtSigner: Signer to create a signed enrollment jwt with. This may be nil if enrollment type is not `jwt`
func NewInitializer(credentialStore oauth_service.ISecureCredentialStore, jwtSigner IJWTSigner) *Initializer {
	return &Initializer{secureCredentialStore: credentialStore, jwtSigner: jwtSigner}
}

func (i *Initializer) InitializeFromFile(ctx context.Context, filename string) (*config.ClientConfig, *management_api_v1.DeviceResponse, error) {
	parser := file_parser.NewFileParser()
	config, err := parser.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	return i.Initialize(ctx, *config)
}

// Initializes the SDK and automatically registers the device if the device has not already been registered
// Returns:
//   - config: config object that should be passed into services when they are initialized (config.Connect must be called before use)
//   - device response: server response from device registration, will be nil on error, or if the device has already been registered
//   - error: Any error that occurs during initialization/registration
func (i *Initializer) Initialize(ctx context.Context, sdkConfig config.SDKInitConfig) (*config.ClientConfig, *management_api_v1.DeviceResponse, error) {

	// Construct the clientConfig
	clientConfig := config.ClientConfig{
		FullyQualifiedDomainName: sdkConfig.FullyQualifiedDomainName,
		IsSecure:                 sdkConfig.IsSecure,
		TenantId:                 sdkConfig.TenantID,
		DeviceId:                 sdkConfig.DeviceID,
	}

	// Check if the device has already been enrolled
	if i.secureCredentialStore.GetClientId() != "" && i.secureCredentialStore.GetClientSecret() != "" {
		return &clientConfig, nil, nil
	}

	// Generate oauth credentials
	oauthCredentials, err := oauth_service.GenerateCredentials()
	if err != nil {
		return &clientConfig, nil, err
	}
	i.secureCredentialStore.SaveCredentials(oauthCredentials.ClientId, oauthCredentials.ClientSecret)

	// Assemble the enrollment credential
	credential := ""
	switch sdkConfig.EnrollmentType {
	case config.None:
		break
	case config.SharedSecret:
		credential = sdkConfig.Credential
	case config.Jwt:
		credential, err = i.jwtSigner.SignJWT(sdkConfig.Credential, sdkConfig.DeviceName, sdkConfig.TenantID, oauthCredentials.ClientId)
		if err != nil {
			return &clientConfig, nil, err
		}
	}

	// Setup the OAuth service
	oauthClientConfig := config.ClientConfig{
		FullyQualifiedDomainName: sdkConfig.FullyQualifiedDomainName,
		IsSecure:                 sdkConfig.IsSecure,
		TenantId:                 sdkConfig.TenantID,
		DeviceId:                 sdkConfig.DeviceID,
	}
	onClose, err := oauthClientConfig.Connect()
	if err != nil {
		return &clientConfig, nil, err
	}
	defer onClose()
	oauthService, err := OauthServiceInit(&oauthClientConfig, i.secureCredentialStore)
	if err != nil {
		return &clientConfig, nil, err
	}

	// Enroll the device
	rsp, err := oauthService.Register(ctx, sdkConfig.DeviceName, credential)
	return &clientConfig, rsp, err
}
