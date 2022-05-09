package token_manager_test

import (
	"context"
	"testing"
	"time"

	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	token_manager "github.com/Sensory-Cloud/go-sdk/pkg/token"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"github.com/google/uuid"
)

type MockOauthService struct {
	ExpiresIn int32
}

func (m *MockOauthService) GetToken(ctx context.Context) (*common.TokenResponse, error) {
	return &common.TokenResponse{
		AccessToken: uuid.NewString(),
		ExpiresIn:   m.ExpiresIn,
		KeyId:       "12345",
		TokenType:   "doesntmatter",
	}, nil
}

func TestGetToken(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthService{100})
	token, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	time.Sleep(10 * time.Millisecond)

	token, err = tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")
}

func TestScheduleTokenRefresh(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthService{0})
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

func TestRefreshToken(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthService{0})
	token, err := tokenManager.GetToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, token != "", "a token should be returned")

	time.Sleep(10 * time.Millisecond)

	newToken, err := tokenManager.RefreshToken()
	test_util.AssertOk(t, err)
	test_util.Assert(t, newToken != token, "a token should be returned")
}

func TestRefreshTokenWithNonExpiredToken(t *testing.T) {
	tokenManager := token_manager.NewTokenManager(&MockOauthService{100})

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
	tokenManagerExp := token_manager.NewTokenManager(&MockOauthService{0})
	_, err := tokenManagerExp.RefreshToken()
	test_util.AssertOk(t, err)

	tokenManagerNotExp := token_manager.NewTokenManager(&MockOauthService{100})
	_, err = tokenManagerNotExp.RefreshToken()
	test_util.AssertOk(t, err)

	test_util.AssertEquals(t, tokenManagerExp.IsTokenExpired(), true)
	test_util.AssertEquals(t, tokenManagerNotExp.IsTokenExpired(), false)
}
