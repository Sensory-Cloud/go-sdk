package token_manager_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Sensory-Cloud/go-sdk/pkg/api/common"
	oauth_api "github.com/Sensory-Cloud/go-sdk/pkg/api/oauth"
	token_manager "github.com/Sensory-Cloud/go-sdk/pkg/auth/token"
	environment "github.com/Sensory-Cloud/go-sdk/pkg/environment"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
)

type MockOauthServiceClient struct {
	ExpiresIn int32
}

func (c *MockOauthServiceClient) GetToken(ctx context.Context, in *oauth_api.TokenRequest, opts ...grpc.CallOption) (*common.TokenResponse, error) {
	return &common.TokenResponse{
		AccessToken: uuid.NewString(),
		ExpiresIn:   c.ExpiresIn,
		KeyId:       "12345",
		TokenType:   "doesntmatter",
	}, nil
}

func (c *MockOauthServiceClient) SignToken(ctx context.Context, in *oauth_api.SignTokenRequest, opts ...grpc.CallOption) (*common.TokenResponse, error) {
	return &common.TokenResponse{
		AccessToken: uuid.NewString(),
		ExpiresIn:   c.ExpiresIn,
		KeyId:       "12345",
		TokenType:   "doesntmatter",
	}, nil
}

func (c *MockOauthServiceClient) GetWhoAmI(ctx context.Context, in *oauth_api.WhoAmIRequest, opts ...grpc.CallOption) (*oauth_api.WhoAmIResponse, error) {
	return &oauth_api.WhoAmIResponse{}, nil
}

func (c *MockOauthServiceClient) GetPublicKey(ctx context.Context, in *oauth_api.PublicKeyRequest, opts ...grpc.CallOption) (*oauth_api.PublicKeyResponse, error) {
	return &oauth_api.PublicKeyResponse{}, nil
}

func TestGetToken(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthServiceClient{100}, "doesn't", "matter")
	token, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	time.Sleep(10 * time.Millisecond)

	token, err = tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")
}

func TestScheduleTokenRefresh(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthServiceClient{0}, "doesn't", "matter")
	token, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	clear := tokenManager.ScheduleTokenRefresh(time.NewTicker(10 * time.Millisecond))
	time.Sleep(20 * time.Millisecond)
	clear()

	newToken, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, newToken != token, "a new token should be returned")
}

func TestSetAuthorizationMetadata(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthServiceClient{100}, "doesn't", "matter")
	ctx, err := tokenManager.SetAuthorizationMetadata(context.Background())
	test_util.AssertOk(t, err)

	metadata := metautils.ExtractOutgoing(ctx)
	token := metadata.Get("authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	cachedToken, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, cachedToken, token)
}

func TestRefreshToken(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthServiceClient{0}, "doesn't", "matter")
	token, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	time.Sleep(10 * time.Millisecond)

	newToken, err := tokenManager.RefreshToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, newToken != token, "a token should be returned")
}

func TestRefreshTokenWithNonExpiredToken(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthServiceClient{100}, "doesn't", "matter")

	token, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	token, err = tokenManager.RefreshToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	newToken, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, newToken, token)
}

func TestIsTokenExpired(t *testing.T) {
	tokenManagerExp := token_manager.NewTokenManager(&MockOauthServiceClient{0}, "doesn't", "matter")
	_, err := tokenManagerExp.RefreshToken()
	test_util.AssertOk(t, err)

	tokenManagerNotExp := token_manager.NewTokenManager(&MockOauthServiceClient{100}, "doesn't", "matter")
	_, err = tokenManagerNotExp.RefreshToken()
	test_util.AssertOk(t, err)

	test_util.AssertEquals(t, tokenManagerExp.IsTokenExpired(), true)
	test_util.AssertEquals(t, tokenManagerNotExp.IsTokenExpired(), false)
}

func TestGetEnvironmentCredentials(t *testing.T) {
	oldToken := environment.Get(environment.AppToken, "")
	defer environment.Set(environment.AppToken, oldToken)

	badTokens := []string{"", "1123asd", "asjhdkj35:asd:3adksjlk"}
	for _, badToken := range badTokens {
		environment.Set(environment.AppToken, badToken)
		clientId, secret, err := token_manager.GetEnvironmentCredentials()
		test_util.AssertEquals(t, clientId, "")
		test_util.AssertEquals(t, secret, "")
		test_util.Assert(t, err != nil, "an error should be returned")
	}

	clientId, secret := "abcd", "1234"
	environment.Set(environment.AppToken, fmt.Sprintf("%s:%s", clientId, secret))
	clientId, secret, err := token_manager.GetEnvironmentCredentials()
	test_util.AssertEquals(t, clientId, clientId)
	test_util.AssertEquals(t, secret, secret)
	test_util.Assert(t, err == nil, "an error should not be returned")
}
