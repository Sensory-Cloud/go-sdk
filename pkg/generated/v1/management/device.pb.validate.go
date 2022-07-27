// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/management/device.proto

package management

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
var _device_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on EnrollDeviceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EnrollDeviceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 127 {
		return EnrollDeviceRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetDeviceId()); l < 1 || l > 127 {
		return EnrollDeviceRequestValidationError{
			field:  "DeviceId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if err := m._validateUuid(m.GetTenantId()); err != nil {
		return EnrollDeviceRequestValidationError{
			field:  "TenantId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if m.GetClient() == nil {
		return EnrollDeviceRequestValidationError{
			field:  "Client",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetClient()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnrollDeviceRequestValidationError{
				field:  "Client",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Credential

	return nil
}

func (m *EnrollDeviceRequest) _validateUuid(uuid string) error {
	if matched := _device_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// EnrollDeviceRequestValidationError is the validation error returned by
// EnrollDeviceRequest.Validate if the designated constraints aren't met.
type EnrollDeviceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrollDeviceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrollDeviceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrollDeviceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrollDeviceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrollDeviceRequestValidationError) ErrorName() string {
	return "EnrollDeviceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e EnrollDeviceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrollDeviceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrollDeviceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrollDeviceRequestValidationError{}

// Validate checks the field values on RenewDeviceCredentialRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RenewDeviceCredentialRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetDeviceId()); l < 1 || l > 127 {
		return RenewDeviceCredentialRequestValidationError{
			field:  "DeviceId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if err := m._validateUuid(m.GetClientId()); err != nil {
		return RenewDeviceCredentialRequestValidationError{
			field:  "ClientId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetTenantId()); err != nil {
		return RenewDeviceCredentialRequestValidationError{
			field:  "TenantId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetCredential()); l < 1 || l > 255 {
		return RenewDeviceCredentialRequestValidationError{
			field:  "Credential",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	return nil
}

func (m *RenewDeviceCredentialRequest) _validateUuid(uuid string) error {
	if matched := _device_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// RenewDeviceCredentialRequestValidationError is the validation error returned
// by RenewDeviceCredentialRequest.Validate if the designated constraints
// aren't met.
type RenewDeviceCredentialRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RenewDeviceCredentialRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RenewDeviceCredentialRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RenewDeviceCredentialRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RenewDeviceCredentialRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RenewDeviceCredentialRequestValidationError) ErrorName() string {
	return "RenewDeviceCredentialRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RenewDeviceCredentialRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRenewDeviceCredentialRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RenewDeviceCredentialRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RenewDeviceCredentialRequestValidationError{}

// Validate checks the field values on DeviceGetWhoAmIRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeviceGetWhoAmIRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeviceGetWhoAmIRequestValidationError is the validation error returned by
// DeviceGetWhoAmIRequest.Validate if the designated constraints aren't met.
type DeviceGetWhoAmIRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceGetWhoAmIRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceGetWhoAmIRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceGetWhoAmIRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceGetWhoAmIRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceGetWhoAmIRequestValidationError) ErrorName() string {
	return "DeviceGetWhoAmIRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeviceGetWhoAmIRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceGetWhoAmIRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceGetWhoAmIRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceGetWhoAmIRequestValidationError{}

// Validate checks the field values on DeviceRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeviceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetDeviceId()); l < 1 || l > 127 {
		return DeviceRequestValidationError{
			field:  "DeviceId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	return nil
}

// DeviceRequestValidationError is the validation error returned by
// DeviceRequest.Validate if the designated constraints aren't met.
type DeviceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceRequestValidationError) ErrorName() string { return "DeviceRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeviceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceRequestValidationError{}

// Validate checks the field values on GetDevicesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetDevicesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TenantId

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDevicesRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetUserId()) > 127 {
		return GetDevicesRequestValidationError{
			field:  "UserId",
			reason: "value length must be at most 127 runes",
		}
	}

	return nil
}

// GetDevicesRequestValidationError is the validation error returned by
// GetDevicesRequest.Validate if the designated constraints aren't met.
type GetDevicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDevicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDevicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDevicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDevicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDevicesRequestValidationError) ErrorName() string {
	return "GetDevicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDevicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDevicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDevicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDevicesRequestValidationError{}

// Validate checks the field values on UpdateDeviceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDeviceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetDeviceId()); l < 1 || l > 127 {
		return UpdateDeviceRequestValidationError{
			field:  "DeviceId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 127 {
		return UpdateDeviceRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	return nil
}

// UpdateDeviceRequestValidationError is the validation error returned by
// UpdateDeviceRequest.Validate if the designated constraints aren't met.
type UpdateDeviceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDeviceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDeviceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDeviceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDeviceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDeviceRequestValidationError) ErrorName() string {
	return "UpdateDeviceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDeviceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDeviceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDeviceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDeviceRequestValidationError{}

// Validate checks the field values on DeviceResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeviceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for DeviceId

	return nil
}

// DeviceResponseValidationError is the validation error returned by
// DeviceResponse.Validate if the designated constraints aren't met.
type DeviceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceResponseValidationError) ErrorName() string { return "DeviceResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeviceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceResponseValidationError{}

// Validate checks the field values on GetDeviceResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetDeviceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for DeviceId

	// no validation rules for UserCount

	return nil
}

// GetDeviceResponseValidationError is the validation error returned by
// GetDeviceResponse.Validate if the designated constraints aren't met.
type GetDeviceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeviceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeviceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeviceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeviceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeviceResponseValidationError) ErrorName() string {
	return "GetDeviceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeviceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeviceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeviceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeviceResponseValidationError{}

// Validate checks the field values on DeviceListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeviceListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetDevices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeviceListResponseValidationError{
					field:  fmt.Sprintf("Devices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceListResponseValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeviceListResponseValidationError is the validation error returned by
// DeviceListResponse.Validate if the designated constraints aren't met.
type DeviceListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceListResponseValidationError) ErrorName() string {
	return "DeviceListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeviceListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceListResponseValidationError{}
