package enrollment_token_manager

import (
	"encoding/binary"
	"errors"
	"time"

	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
)

var (
	tokenSuffix      = "-token"
	expirationSuffix = "-expiration"
)

// ISecureTokenStore is a interface for securely storing arbitrary data
type ISecureTokenStore interface {
	// Saves a new byte array under a specific key value
	SaveData(string, []byte) error
	// Retrieves a byte array for a specific key value. `nil, nil` should be returned if the data was not found
	GetData(string) ([]byte, error)
	// Deletes a byte array for a specific key value. `nil` should be returned if the data to be deleted was not found
	DeleteData(string) error
}

// IEnrollmentTokenManager is an interface to manage and store Sensory Cloud enrollment tokens
type IEnrollmentTokenManager interface {
	// Saves a new enrollment token for a given enrollment ID.
	SaveEnrollmentToken(string, *common.EnrollmentToken) error
	// Retrieves the enrollment token for an enrollment ID, `nil, nil` is returned if no token was found for the given enrollment ID
	GetEnrollmentToken(string) ([]byte, error)
	// Deletes a saved enrollment token, `nil` is returned if there was no token saved for the given enrollment ID
	DeleteEnrollmentToken(string) error
}

// EnrollmentTokenManager saves Sensory Cloud enrollment tokens and manages their expirations
type EnrollmentTokenManager struct {
	tokenStore ISecureTokenStore
}

// Returns a new EnrollmentTokenManager
func NewEnrollmentTokenManager(tokenStore ISecureTokenStore) *EnrollmentTokenManager {
	return &EnrollmentTokenManager{tokenStore: tokenStore}
}

// Saves a new enrollment token for a given enrollment ID.
func (t *EnrollmentTokenManager) SaveEnrollmentToken(enrollmentID string, enrollmentToken *common.EnrollmentToken) error {
	err := t.DeleteEnrollmentToken(enrollmentID)
	if err != nil {
		return err
	}

	err = t.tokenStore.SaveData(enrollmentID+tokenSuffix, enrollmentToken.GetToken())
	if err != nil {
		return err
	}

	if enrollmentToken.Expiration > 0 {
		expiration := time.Now().Unix() + enrollmentToken.Expiration
		expData := make([]byte, 8)
		binary.LittleEndian.PutUint64(expData, uint64(expiration))
		err = t.tokenStore.SaveData(enrollmentID+expirationSuffix, expData)
		if err != nil {
			return err
		}
	}

	return nil
}

// Retrieves the enrollment token for an enrollment ID, `nil, nil` is returned if no token was found for the given enrollment ID
func (t *EnrollmentTokenManager) GetEnrollmentToken(enrollmentID string) ([]byte, error) {
	token, err := t.tokenStore.GetData(enrollmentID + tokenSuffix)
	if err != nil {
		return nil, err
	} else if token == nil {
		// Return nil on token not found instead of throwing an error
		return nil, nil
	}

	expData, err := t.tokenStore.GetData(enrollmentID + expirationSuffix)
	if err != nil {
		return nil, err
	} else if expData == nil {
		// Return the token when it has no expiration date
		return token, nil
	}

	expiration := int64(binary.LittleEndian.Uint64(expData))
	now := time.Now().Unix()
	if now > expiration {
		// The enrollment token has expired
		err = t.tokenStore.DeleteData(enrollmentID)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("the requested enrollment token has expired")
	}

	return token, nil
}

// Deletes a saved enrollment token, `nil` is returned if there was no token saved for the given enrollment ID
func (t *EnrollmentTokenManager) DeleteEnrollmentToken(enrollmentID string) error {
	err := t.tokenStore.DeleteData(enrollmentID + tokenSuffix)
	if err != nil {
		return err
	}

	return t.tokenStore.DeleteData(enrollmentID + expirationSuffix)
}
