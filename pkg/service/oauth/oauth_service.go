package oauth_service

import (
	"fmt"

	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	oauth "github.com/Sensory-Cloud/go-sdk/pkg/generated/oauth"
	management_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/management"
	crypto_service "github.com/Sensory-Cloud/go-sdk/pkg/service/crypto"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type ISecureCredentialStore interface {
	GetClientId() string
	GetClientSecret() string
}

type OauthService struct {
	config                *config.ClientConfig
	oauthClient           oauth.OauthServiceClient
	deviceClient          management_api_v1.DeviceServiceClient
	secureCredentialStore ISecureCredentialStore
}

type OauthClient struct {
	ClientId     string
	ClientSecret string
}

// NewOauthService creates a new service to handle OAuth authentication with Sensory Cloud
func NewOauthService(config *config.ClientConfig, secureCredentialStore ISecureCredentialStore) (*OauthService, error) {
	client, err := config.GetClient()
	if err != nil {
		return nil, err
	}

	oauthClient := oauth.NewOauthServiceClient(client)
	deviceClient := management_api_v1.NewDeviceServiceClient(client)
	return &OauthService{config, oauthClient, deviceClient, secureCredentialStore}, nil
}

// Can be called to generate secure and guaranteed unique credentials.
// Should be used the first time the SDK registers an OAuth token with the cloud.
func GenerateCredentials() (*OauthClient, error) {
	clientSecret, err := crypto_service.GenerateRandomString(24)
	if err != nil {
		return nil, err
	}

	return &OauthClient{uuid.NewString(), clientSecret}, nil
}

// Get information about the current registered device as inferred by the OAuth credentials supplied by the credential manager.
// A new token is request every time this call is made, so use sparingly.
func (s *OauthService) GetWhoAmI() (*management_api_v1.DeviceResponse, error) {
	token, err := s.GetToken()
	if err != nil {
		return nil, err
	}

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %s", token))

	return s.deviceClient.GetWhoAmI(ctx, &management_api_v1.DeviceGetWhoAmIRequest{})
}

// GetToken retrieves a token from Sensory Cloud based on the configured credentials
func (s *OauthService) GetToken() (*common.TokenResponse, error) {
	clientId := s.secureCredentialStore.GetClientId()
	if clientId == "" {
		return nil, error_util.Errorf(codes.InvalidArgument, "null clientId was returned from the secure credential store")
	}

	clientSecret := s.secureCredentialStore.GetClientSecret()
	if clientSecret == "" {
		return nil, error_util.Errorf(codes.InvalidArgument, "null clientSecret was returned from the secure credential store")
	}

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	return s.oauthClient.GetToken(ctx, &oauth.TokenRequest{ClientId: clientId, Secret: clientSecret})
}

// Register credentials provided by the attached SecureCredentialStore to Sensory Cloud. This function should only be called
// once per unique credential pair. An error will be thrown if registration fails.
func (s *OauthService) Register(deviceName string, credential string) (*management_api_v1.DeviceResponse, error) {
	clientId := s.secureCredentialStore.GetClientId()
	if clientId == "" {
		return nil, error_util.Errorf(codes.InvalidArgument, "null clientId was returned from the secure credential store")
	}

	clientSecret := s.secureCredentialStore.GetClientSecret()
	if clientSecret == "" {
		return nil, error_util.Errorf(codes.InvalidArgument, "null clientSecret was returned from the secure credential store")
	}

	client := common.GenericClient{ClientId: clientId, Secret: clientSecret}
	request := management_api_v1.EnrollDeviceRequest{
		Client:     &client,
		DeviceId:   s.config.DeviceId,
		Name:       deviceName,
		Credential: credential,
		TenantId:   s.config.TenantId,
	}

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	return s.deviceClient.EnrollDevice(ctx, &request)
}

// RenewDeviceCredential allows a client to re-register with Sensory Cloud if their credentials have expired
func (s *OauthService) RenewDeviceCredential(credential string) (*management_api_v1.DeviceResponse, error) {
	clientId := s.secureCredentialStore.GetClientId()
	if clientId == "" {
		return nil, error_util.Errorf(codes.InvalidArgument, "null clientId was returned from the secure credential store")
	}

	request := management_api_v1.RenewDeviceCredentialRequest{
		ClientId:   clientId,
		Credential: credential,
		DeviceId:   s.config.DeviceId,
		TenantId:   s.config.TenantId,
	}

	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	return s.deviceClient.RenewDeviceCredential(ctx, &request)
}