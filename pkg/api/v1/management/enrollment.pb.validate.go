// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/management/enrollment.proto

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

	common "github.com/Sensory-Cloud/go-sdk/pkg/api/common"
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

	_ = common.ModelType(0)

	_ = common.ModelType(0)
)

// define the regex for a UUID once up-front
var _enrollment_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetEnrollmentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEnrollmentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetUserId()); l < 1 || l > 127 {
		return GetEnrollmentsRequestValidationError{
			field:  "UserId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	return nil
}

// GetEnrollmentsRequestValidationError is the validation error returned by
// GetEnrollmentsRequest.Validate if the designated constraints aren't met.
type GetEnrollmentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEnrollmentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEnrollmentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEnrollmentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEnrollmentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEnrollmentsRequestValidationError) ErrorName() string {
	return "GetEnrollmentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEnrollmentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEnrollmentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEnrollmentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEnrollmentsRequestValidationError{}

// Validate checks the field values on GetEnrollmentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEnrollmentsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEnrollments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEnrollmentsResponseValidationError{
					field:  fmt.Sprintf("Enrollments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for IsRequestorTrusted

	return nil
}

// GetEnrollmentsResponseValidationError is the validation error returned by
// GetEnrollmentsResponse.Validate if the designated constraints aren't met.
type GetEnrollmentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEnrollmentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEnrollmentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEnrollmentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEnrollmentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEnrollmentsResponseValidationError) ErrorName() string {
	return "GetEnrollmentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEnrollmentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEnrollmentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEnrollmentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEnrollmentsResponseValidationError{}

// Validate checks the field values on EnrollmentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EnrollmentResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnrollmentResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnrollmentResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Description

	// no validation rules for ModelName

	// no validation rules for ModelType

	// no validation rules for ModelVersion

	// no validation rules for DeviceId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCompression()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnrollmentResponseValidationError{
				field:  "Compression",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DeviceName

	// no validation rules for DidEnrollWithLiveness

	return nil
}

// EnrollmentResponseValidationError is the validation error returned by
// EnrollmentResponse.Validate if the designated constraints aren't met.
type EnrollmentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrollmentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrollmentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrollmentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrollmentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrollmentResponseValidationError) ErrorName() string {
	return "EnrollmentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EnrollmentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrollmentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrollmentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrollmentResponseValidationError{}

// Validate checks the field values on GetEnrollmentGroupsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEnrollmentGroupsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEnrollmentGroups() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEnrollmentGroupsResponseValidationError{
					field:  fmt.Sprintf("EnrollmentGroups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetEnrollmentGroupsResponseValidationError is the validation error returned
// by GetEnrollmentGroupsResponse.Validate if the designated constraints
// aren't met.
type GetEnrollmentGroupsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEnrollmentGroupsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEnrollmentGroupsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEnrollmentGroupsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEnrollmentGroupsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEnrollmentGroupsResponseValidationError) ErrorName() string {
	return "GetEnrollmentGroupsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEnrollmentGroupsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEnrollmentGroupsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEnrollmentGroupsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEnrollmentGroupsResponseValidationError{}

// Validate checks the field values on EnrollmentGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EnrollmentGroupResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnrollmentGroupResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnrollmentGroupResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for ModelName

	// no validation rules for ModelType

	// no validation rules for ModelVersion

	// no validation rules for UserId

	for idx, item := range m.GetEnrollments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnrollmentGroupResponseValidationError{
					field:  fmt.Sprintf("Enrollments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EnrollmentGroupResponseValidationError is the validation error returned by
// EnrollmentGroupResponse.Validate if the designated constraints aren't met.
type EnrollmentGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrollmentGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrollmentGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrollmentGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrollmentGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrollmentGroupResponseValidationError) ErrorName() string {
	return "EnrollmentGroupResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EnrollmentGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrollmentGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrollmentGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrollmentGroupResponseValidationError{}

// Validate checks the field values on CreateEnrollmentGroupRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEnrollmentGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetId()); l < 1 || l > 127 {
		return CreateEnrollmentGroupRequestValidationError{
			field:  "Id",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 255 {
		return CreateEnrollmentGroupRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) > 1023 {
		return CreateEnrollmentGroupRequestValidationError{
			field:  "Description",
			reason: "value length must be at most 1023 runes",
		}
	}

	if l := utf8.RuneCountInString(m.GetModelName()); l < 1 || l > 255 {
		return CreateEnrollmentGroupRequestValidationError{
			field:  "ModelName",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetUserId()); l < 1 || l > 127 {
		return CreateEnrollmentGroupRequestValidationError{
			field:  "UserId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	return nil
}

// CreateEnrollmentGroupRequestValidationError is the validation error returned
// by CreateEnrollmentGroupRequest.Validate if the designated constraints
// aren't met.
type CreateEnrollmentGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEnrollmentGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEnrollmentGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEnrollmentGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEnrollmentGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEnrollmentGroupRequestValidationError) ErrorName() string {
	return "CreateEnrollmentGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEnrollmentGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEnrollmentGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEnrollmentGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEnrollmentGroupRequestValidationError{}

// Validate checks the field values on AppendEnrollmentGroupRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AppendEnrollmentGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetGroupId()); l < 1 || l > 127 {
		return AppendEnrollmentGroupRequestValidationError{
			field:  "GroupId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	return nil
}

// AppendEnrollmentGroupRequestValidationError is the validation error returned
// by AppendEnrollmentGroupRequest.Validate if the designated constraints
// aren't met.
type AppendEnrollmentGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppendEnrollmentGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppendEnrollmentGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppendEnrollmentGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppendEnrollmentGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppendEnrollmentGroupRequestValidationError) ErrorName() string {
	return "AppendEnrollmentGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AppendEnrollmentGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppendEnrollmentGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppendEnrollmentGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppendEnrollmentGroupRequestValidationError{}

// Validate checks the field values on DeleteEnrollmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteEnrollmentRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return DeleteEnrollmentRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *DeleteEnrollmentRequest) _validateUuid(uuid string) error {
	if matched := _enrollment_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteEnrollmentRequestValidationError is the validation error returned by
// DeleteEnrollmentRequest.Validate if the designated constraints aren't met.
type DeleteEnrollmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEnrollmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEnrollmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEnrollmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEnrollmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEnrollmentRequestValidationError) ErrorName() string {
	return "DeleteEnrollmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEnrollmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEnrollmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEnrollmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEnrollmentRequestValidationError{}

// Validate checks the field values on DeleteEnrollmentGroupRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteEnrollmentGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetId()); l < 1 || l > 127 {
		return DeleteEnrollmentGroupRequestValidationError{
			field:  "Id",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	return nil
}

// DeleteEnrollmentGroupRequestValidationError is the validation error returned
// by DeleteEnrollmentGroupRequest.Validate if the designated constraints
// aren't met.
type DeleteEnrollmentGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEnrollmentGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEnrollmentGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEnrollmentGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEnrollmentGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEnrollmentGroupRequestValidationError) ErrorName() string {
	return "DeleteEnrollmentGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEnrollmentGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEnrollmentGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEnrollmentGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEnrollmentGroupRequestValidationError{}
