package enrollment_token_manager_test

import (
	"encoding/binary"
	"testing"
	"time"

	enrollment_token_manager "github.com/Sensory-Cloud/go-sdk/pkg/enrollment_token"
	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

var (
	mockTokenExpires = &common.EnrollmentToken{Token: []byte("Mock enrollment token"), Expiration: 300}
	mockToken        = &common.EnrollmentToken{Token: []byte("Some token value"), Expiration: 0}
)

func TestSaveGetToken(t *testing.T) {
	mockPersistence := NewMockSecureTokenStore()
	tokenManager := enrollment_token_manager.NewEnrollmentTokenManager(mockPersistence)

	res, err := tokenManager.GetEnrollmentToken("bogus")
	test_util.AssertOk(t, err)
	test_util.Assert(t, res == nil, "Nil should be returned on token not found")

	err = tokenManager.SaveEnrollmentToken("enrollment", mockToken)
	test_util.AssertOk(t, err)

	res, err = tokenManager.GetEnrollmentToken("enrollment")
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, res, mockToken.Token)

	err = tokenManager.SaveEnrollmentToken("enrollment", mockTokenExpires)
	test_util.AssertOk(t, err)

	res, err = tokenManager.GetEnrollmentToken("enrollment")
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, res, mockTokenExpires.Token)

	mockExpiration := time.Now().Unix() - 300
	expData := make([]byte, 8)
	binary.LittleEndian.PutUint64(expData, uint64(mockExpiration))
	err = mockPersistence.SaveData("enrollment-expiration", expData)
	test_util.AssertOk(t, err)

	res, err = tokenManager.GetEnrollmentToken("enrollment")
	test_util.Assert(t, err != nil, "Error should be thrown on expired token")
	test_util.Assert(t, res == nil, "Expired token should not be returned")
}

func TestDeleteToken(t *testing.T) {
	mockPersistence := NewMockSecureTokenStore()
	tokenManager := enrollment_token_manager.NewEnrollmentTokenManager(mockPersistence)

	err := tokenManager.DeleteEnrollmentToken("bogus")
	test_util.Assert(t, err == nil, "Nil should be returned on token not found")

	err = tokenManager.SaveEnrollmentToken("enrollment", mockToken)
	test_util.AssertOk(t, err)

	err = tokenManager.DeleteEnrollmentToken("enrollment")
	test_util.AssertOk(t, err)

	res, err := tokenManager.GetEnrollmentToken("enrollment")
	test_util.AssertOk(t, err)
	test_util.Assert(t, res == nil, "Deleted token should not be returned")
}

type MockSecureTokenStore struct {
	store map[string][]byte
	Err   error
}

func NewMockSecureTokenStore() *MockSecureTokenStore {
	return &MockSecureTokenStore{make(map[string][]byte, 0), nil}
}

func (m *MockSecureTokenStore) SaveData(key string, value []byte) error {
	if m.Err != nil {
		return m.Err
	}

	m.store[key] = value
	return nil
}

func (m *MockSecureTokenStore) GetData(key string) ([]byte, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	value, present := m.store[key]
	if !present {
		return nil, nil
	}
	return value, nil
}

func (m *MockSecureTokenStore) DeleteData(key string) error {
	if m.Err != nil {
		return m.Err
	}

	delete(m.store, key)
	return nil
}
