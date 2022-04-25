package token_manager

import (
	"context"
	"fmt"
	"strings"
	"time"

	common "github.com/Sensory-Cloud/go-sdk/pkg/api/common"
	oauth_api "github.com/Sensory-Cloud/go-sdk/pkg/api/oauth"
	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	environment "github.com/Sensory-Cloud/go-sdk/pkg/environment"
	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	log_util "github.com/Sensory-Cloud/go-sdk/pkg/util/log"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
)

var (
	log, _                            = log_util.GetSugaredLogger()
	headerAuthorize                   = "authorization"
	tokenExpiresBufferPercent float32 = 0.5
)

// IAuthorizationMetadataService is an interface to manage authorization metadata on contexts
type IAuthorizationMetadataService interface {
	// SetAuthorizationMetadata appends authorization metadata to a context
	SetAuthorizationMetadata(ctx context.Context) (context.Context, error)
}

// ITokenManager is an interface representing any token manager
type ITokenManager interface {
	IAuthorizationMetadataService
	// GetToken returns the currently cached token
	GetToken() (string, error)
	// RefreshToken obtains and caches a new token
	RefreshToken() error
	// IsTokenExpired returns true if the token is expired
	IsTokenExpired() bool
}

// TokenManager retrieves and caches OAuth tokens
type TokenManager struct {
	oauthServiceClient oauth_api.OauthServiceClient
	clientId           string
	secret             string
	token              string
	keyId              string
	expires            time.Time
}

// NewAppTokenManager returns a new TokenManger using the APP_TOKEN environment variable
func NewAppTokenManager(oauthServiceClient oauth_api.OauthServiceClient) (*TokenManager, error) {
	clientId, secret, err := GetEnvironmentCredentials()
	if err != nil {
		return nil, err
	}

	return NewTokenManager(oauthServiceClient, clientId, secret), nil
}

// NewTokenManager returns a new TokenManager.
// It is recommended to call tokenManager.ScheduleTokenRefresh() after construction
// in order to gain the full advantages of the TokenManager.
//
// This function returns a clear function that can be used to stop the scheduled task.
func NewTokenManager(oauthServiceClient oauth_api.OauthServiceClient, clientId string, secret string) *TokenManager {
	tokenManager := &TokenManager{oauthServiceClient: oauthServiceClient, clientId: clientId, secret: secret}
	return tokenManager
}

// ScheduleTokenRefresh checks and refreshes the token at a fixed interval specified by the ticker on a background thread.
// Calls to this function are non-blocking, as the function spins up a goroutine to handle the background check.
func (t *TokenManager) ScheduleTokenRefresh(tokenRefreshInterval *time.Ticker) (clear func()) {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-tokenRefreshInterval.C:
				_, err := t.RefreshToken()
				if err != nil {
					log.Errorf("Failed to refresh access token: %s", err.Error())
				}
			case <-done:
				log.Info("Stopping key token refresh")
				return
			}
		}
	}()

	return func() { done <- true }
}

// GetToken returns the cached token.
// If a token is not yet retrieved, GetToken() will call RefreshToken()
func (t *TokenManager) GetToken() (string, error) {
	if t.token == "" {
		return t.RefreshToken()
	}
	return t.token, nil
}

func (t *TokenManager) SetAuthorizationMetadata(ctx context.Context) (context.Context, error) {
	token, err := t.GetToken()
	if err != nil {
		return nil, err
	}

	return metadata.AppendToOutgoingContext(ctx, headerAuthorize, fmt.Sprintf("Bearer %s", token)), nil
}

// RefreshToken fetches an caches a new token on if the current token is expired
func (t *TokenManager) RefreshToken() (string, error) {
	if !t.IsTokenExpired() {
		return t.token, nil
	}

	token, err := t.fetchToken()
	if err != nil {
		return t.token, err
	}

	expiresWithBuffer := token.ExpiresIn - int32(float32(token.ExpiresIn)*tokenExpiresBufferPercent)

	t.token = token.AccessToken
	t.keyId = token.KeyId
	t.expires = time.Now().UTC().Add(time.Duration(expiresWithBuffer) * time.Second)
	return t.token, nil
}

// IsTokenExpired returns true if the current token is not set or expired
func (t *TokenManager) IsTokenExpired() bool {
	return t.token == "" || t.expires.IsZero() || time.Now().UTC().After(t.expires.UTC())
}

func GetEnvironmentCredentials() (clientId string, secret string, err error) {
	rawToken := environment.Get(environment.AppToken, "")
	if rawToken == "" {
		return "", "", error_util.Error(codes.FailedPrecondition, "config error: APP_TOKEN is empty!")
	}

	splitToken := strings.Split(rawToken, ":")
	if len(splitToken) != 2 {
		return "", "", error_util.Error(codes.FailedPrecondition, "config error: APP_TOKEN is incorrectly formatted")
	}

	return splitToken[0], splitToken[1], nil
}

func (t *TokenManager) fetchToken() (*common.TokenResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	return t.oauthServiceClient.GetToken(ctx, &oauth_api.TokenRequest{ClientId: t.clientId, Secret: t.secret})
}
