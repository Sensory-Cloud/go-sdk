// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/common.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _common_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CompressionConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CompressionConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CompressionConfigurationValidationError is the validation error returned by
// CompressionConfiguration.Validate if the designated constraints aren't met.
type CompressionConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompressionConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompressionConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompressionConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompressionConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompressionConfigurationValidationError) ErrorName() string {
	return "CompressionConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e CompressionConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompressionConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompressionConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompressionConfigurationValidationError{}

// Validate checks the field values on TokenResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TokenResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for ExpiresIn

	// no validation rules for KeyId

	// no validation rules for TokenType

	return nil
}

// TokenResponseValidationError is the validation error returned by
// TokenResponse.Validate if the designated constraints aren't met.
type TokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenResponseValidationError) ErrorName() string { return "TokenResponseValidationError" }

// Error satisfies the builtin error interface
func (e TokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenResponseValidationError{}

// Validate checks the field values on ServiceHealth with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ServiceHealth) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsHealthy

	// no validation rules for Name

	// no validation rules for Message

	return nil
}

// ServiceHealthValidationError is the validation error returned by
// ServiceHealth.Validate if the designated constraints aren't met.
type ServiceHealthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceHealthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceHealthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceHealthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceHealthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceHealthValidationError) ErrorName() string { return "ServiceHealthValidationError" }

// Error satisfies the builtin error interface
func (e ServiceHealthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceHealth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceHealthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceHealthValidationError{}

// Validate checks the field values on ServerHealthResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServerHealthResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsHealthy

	// no validation rules for ServerVersion

	// no validation rules for Id

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerHealthResponseValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ServerType

	// no validation rules for IsLeader

	return nil
}

// ServerHealthResponseValidationError is the validation error returned by
// ServerHealthResponse.Validate if the designated constraints aren't met.
type ServerHealthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerHealthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerHealthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerHealthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerHealthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerHealthResponseValidationError) ErrorName() string {
	return "ServerHealthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ServerHealthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerHealthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerHealthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerHealthResponseValidationError{}

// Validate checks the field values on SystemSummary with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SystemSummary) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCpu() == nil {
		return SystemSummaryValidationError{
			field:  "Cpu",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetCpu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SystemSummaryValidationError{
				field:  "Cpu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetMemory() == nil {
		return SystemSummaryValidationError{
			field:  "Memory",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMemory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SystemSummaryValidationError{
				field:  "Memory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SystemSummaryValidationError is the validation error returned by
// SystemSummary.Validate if the designated constraints aren't met.
type SystemSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SystemSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SystemSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SystemSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SystemSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SystemSummaryValidationError) ErrorName() string { return "SystemSummaryValidationError" }

// Error satisfies the builtin error interface
func (e SystemSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSystemSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SystemSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SystemSummaryValidationError{}

// Validate checks the field values on CpuSummary with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CpuSummary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for User

	// no validation rules for Nice

	// no validation rules for System

	// no validation rules for Idle

	// no validation rules for IoWait

	// no validation rules for Irq

	// no validation rules for SoftIrq

	// no validation rules for Steal

	// no validation rules for Guest

	// no validation rules for GuestNice

	return nil
}

// CpuSummaryValidationError is the validation error returned by
// CpuSummary.Validate if the designated constraints aren't met.
type CpuSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CpuSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CpuSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CpuSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CpuSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CpuSummaryValidationError) ErrorName() string { return "CpuSummaryValidationError" }

// Error satisfies the builtin error interface
func (e CpuSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCpuSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CpuSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CpuSummaryValidationError{}

// Validate checks the field values on MemorySummary with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MemorySummary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MemTotal

	// no validation rules for MemFree

	// no validation rules for MemAvailable

	return nil
}

// MemorySummaryValidationError is the validation error returned by
// MemorySummary.Validate if the designated constraints aren't met.
type MemorySummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemorySummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemorySummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemorySummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemorySummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemorySummaryValidationError) ErrorName() string { return "MemorySummaryValidationError" }

// Error satisfies the builtin error interface
func (e MemorySummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemorySummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemorySummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemorySummaryValidationError{}

// Validate checks the field values on GenericClient with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GenericClient) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetClientId()); err != nil {
		return GenericClientValidationError{
			field:  "ClientId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetSecret()) < 10 {
		return GenericClientValidationError{
			field:  "Secret",
			reason: "value length must be at least 10 runes",
		}
	}

	return nil
}

func (m *GenericClient) _validateUuid(uuid string) error {
	if matched := _common_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GenericClientValidationError is the validation error returned by
// GenericClient.Validate if the designated constraints aren't met.
type GenericClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenericClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenericClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenericClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenericClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenericClientValidationError) ErrorName() string { return "GenericClientValidationError" }

// Error satisfies the builtin error interface
func (e GenericClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenericClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenericClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenericClientValidationError{}

// Validate checks the field values on TenantResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TenantResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TenantResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TenantResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TenantResponseValidationError is the validation error returned by
// TenantResponse.Validate if the designated constraints aren't met.
type TenantResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantResponseValidationError) ErrorName() string { return "TenantResponseValidationError" }

// Error satisfies the builtin error interface
func (e TenantResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenantResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantResponseValidationError{}

// Validate checks the field values on PaginationOptions with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PaginationOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Ordering

	// no validation rules for Decending

	// no validation rules for PageIndex

	// no validation rules for PageSize

	return nil
}

// PaginationOptionsValidationError is the validation error returned by
// PaginationOptions.Validate if the designated constraints aren't met.
type PaginationOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationOptionsValidationError) ErrorName() string {
	return "PaginationOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e PaginationOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginationOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationOptionsValidationError{}

// Validate checks the field values on PaginationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PaginationResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Ordering

	// no validation rules for Decending

	// no validation rules for TotalCount

	// no validation rules for PageSize

	// no validation rules for PrevPageIndex

	// no validation rules for CurrentPageIndex

	// no validation rules for NextPageIndex

	return nil
}

// PaginationResponseValidationError is the validation error returned by
// PaginationResponse.Validate if the designated constraints aren't met.
type PaginationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationResponseValidationError) ErrorName() string {
	return "PaginationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PaginationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationResponseValidationError{}

// Validate checks the field values on EnrollmentToken with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *EnrollmentToken) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for Expiration

	return nil
}

// EnrollmentTokenValidationError is the validation error returned by
// EnrollmentToken.Validate if the designated constraints aren't met.
type EnrollmentTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrollmentTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrollmentTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrollmentTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrollmentTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrollmentTokenValidationError) ErrorName() string { return "EnrollmentTokenValidationError" }

// Error satisfies the builtin error interface
func (e EnrollmentTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrollmentToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrollmentTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrollmentTokenValidationError{}

// Validate checks the field values on CreateKeyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateKeyRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 127 {
		return CreateKeyRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if _, ok := KeyType_name[int32(m.GetKeyType())]; !ok {
		return CreateKeyRequestValidationError{
			field:  "KeyType",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for Value

	// no validation rules for Expiration

	return nil
}

// CreateKeyRequestValidationError is the validation error returned by
// CreateKeyRequest.Validate if the designated constraints aren't met.
type CreateKeyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateKeyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateKeyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateKeyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateKeyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateKeyRequestValidationError) ErrorName() string { return "CreateKeyRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateKeyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateKeyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateKeyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateKeyRequestValidationError{}

// Validate checks the field values on KeyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *KeyResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for KeyType

	// no validation rules for Expiration

	// no validation rules for TenantId

	// no validation rules for Disabled

	return nil
}

// KeyResponseValidationError is the validation error returned by
// KeyResponse.Validate if the designated constraints aren't met.
type KeyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyResponseValidationError) ErrorName() string { return "KeyResponseValidationError" }

// Error satisfies the builtin error interface
func (e KeyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyResponseValidationError{}
