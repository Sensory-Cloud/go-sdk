// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/video/video.proto

package video

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

	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
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

	_ = common.TechnologyType(0)
)

// define the regex for a UUID once up-front
var _video_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on VideoModel with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *VideoModel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for IsEnrollable

	// no validation rules for ModelType

	// no validation rules for FixedObject

	// no validation rules for Technology

	// no validation rules for IsLivenessSupported

	return nil
}

// VideoModelValidationError is the validation error returned by
// VideoModel.Validate if the designated constraints aren't met.
type VideoModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoModelValidationError) ErrorName() string { return "VideoModelValidationError" }

// Error satisfies the builtin error interface
func (e VideoModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideoModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoModelValidationError{}

// Validate checks the field values on GetModelsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetModelsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetModelsRequestValidationError is the validation error returned by
// GetModelsRequest.Validate if the designated constraints aren't met.
type GetModelsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetModelsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetModelsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetModelsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetModelsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetModelsRequestValidationError) ErrorName() string { return "GetModelsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetModelsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetModelsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetModelsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetModelsRequestValidationError{}

// Validate checks the field values on GetModelsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetModelsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetModels() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetModelsResponseValidationError{
					field:  fmt.Sprintf("Models[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetModelsResponseValidationError is the validation error returned by
// GetModelsResponse.Validate if the designated constraints aren't met.
type GetModelsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetModelsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetModelsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetModelsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetModelsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetModelsResponseValidationError) ErrorName() string {
	return "GetModelsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetModelsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetModelsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetModelsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetModelsResponseValidationError{}

// Validate checks the field values on CreateEnrollmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEnrollmentRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.StreamingRequest.(type) {

	case *CreateEnrollmentRequest_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateEnrollmentRequestValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CreateEnrollmentRequest_ImageContent:
		// no validation rules for ImageContent

	default:
		return CreateEnrollmentRequestValidationError{
			field:  "StreamingRequest",
			reason: "value is required",
		}

	}

	return nil
}

// CreateEnrollmentRequestValidationError is the validation error returned by
// CreateEnrollmentRequest.Validate if the designated constraints aren't met.
type CreateEnrollmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEnrollmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEnrollmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEnrollmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEnrollmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEnrollmentRequestValidationError) ErrorName() string {
	return "CreateEnrollmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEnrollmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEnrollmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEnrollmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEnrollmentRequestValidationError{}

// Validate checks the field values on AuthenticateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthenticateRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.StreamingRequest.(type) {

	case *AuthenticateRequest_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthenticateRequestValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AuthenticateRequest_ImageContent:
		// no validation rules for ImageContent

	default:
		return AuthenticateRequestValidationError{
			field:  "StreamingRequest",
			reason: "value is required",
		}

	}

	return nil
}

// AuthenticateRequestValidationError is the validation error returned by
// AuthenticateRequest.Validate if the designated constraints aren't met.
type AuthenticateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticateRequestValidationError) ErrorName() string {
	return "AuthenticateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticateRequestValidationError{}

// Validate checks the field values on ValidateRecognitionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ValidateRecognitionRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.StreamingRequest.(type) {

	case *ValidateRecognitionRequest_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValidateRecognitionRequestValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ValidateRecognitionRequest_ImageContent:
		// no validation rules for ImageContent

	default:
		return ValidateRecognitionRequestValidationError{
			field:  "StreamingRequest",
			reason: "value is required",
		}

	}

	return nil
}

// ValidateRecognitionRequestValidationError is the validation error returned
// by ValidateRecognitionRequest.Validate if the designated constraints aren't met.
type ValidateRecognitionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateRecognitionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateRecognitionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateRecognitionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateRecognitionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateRecognitionRequestValidationError) ErrorName() string {
	return "ValidateRecognitionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateRecognitionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateRecognitionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateRecognitionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateRecognitionRequestValidationError{}

// Validate checks the field values on CreateEnrollmentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEnrollmentResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PercentComplete

	// no validation rules for IsAlive

	// no validation rules for EnrollmentId

	// no validation rules for ModelName

	// no validation rules for ModelVersion

	// no validation rules for Score

	if v, ok := interface{}(m.GetEnrollmentToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEnrollmentResponseValidationError{
				field:  "EnrollmentToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DidFindFace

	// no validation rules for ProbabilityFace

	return nil
}

// CreateEnrollmentResponseValidationError is the validation error returned by
// CreateEnrollmentResponse.Validate if the designated constraints aren't met.
type CreateEnrollmentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEnrollmentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEnrollmentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEnrollmentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEnrollmentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEnrollmentResponseValidationError) ErrorName() string {
	return "CreateEnrollmentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEnrollmentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEnrollmentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEnrollmentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEnrollmentResponseValidationError{}

// Validate checks the field values on AuthenticateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthenticateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Score

	// no validation rules for IsAlive

	if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthenticateResponseValidationError{
				field:  "Token",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for EnrollmentId

	// no validation rules for DidFindFace

	// no validation rules for ProbabilityFace

	return nil
}

// AuthenticateResponseValidationError is the validation error returned by
// AuthenticateResponse.Validate if the designated constraints aren't met.
type AuthenticateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticateResponseValidationError) ErrorName() string {
	return "AuthenticateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticateResponseValidationError{}

// Validate checks the field values on LivenessRecognitionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LivenessRecognitionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsAlive

	// no validation rules for Score

	// no validation rules for DidFindFace

	// no validation rules for ProbabilityFace

	return nil
}

// LivenessRecognitionResponseValidationError is the validation error returned
// by LivenessRecognitionResponse.Validate if the designated constraints
// aren't met.
type LivenessRecognitionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LivenessRecognitionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LivenessRecognitionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LivenessRecognitionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LivenessRecognitionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LivenessRecognitionResponseValidationError) ErrorName() string {
	return "LivenessRecognitionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LivenessRecognitionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLivenessRecognitionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LivenessRecognitionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LivenessRecognitionResponseValidationError{}

// Validate checks the field values on CreateEnrollmentConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEnrollmentConfig) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetUserId()); l < 1 || l > 127 {
		return CreateEnrollmentConfigValidationError{
			field:  "UserId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetDeviceId()); l < 1 || l > 127 {
		return CreateEnrollmentConfigValidationError{
			field:  "DeviceId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetModelName()); l < 1 || l > 255 {
		return CreateEnrollmentConfigValidationError{
			field:  "ModelName",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) > 1023 {
		return CreateEnrollmentConfigValidationError{
			field:  "Description",
			reason: "value length must be at most 1023 runes",
		}
	}

	// no validation rules for IsLivenessEnabled

	if _, ok := RecognitionThreshold_name[int32(m.GetLivenessThreshold())]; !ok {
		return CreateEnrollmentConfigValidationError{
			field:  "LivenessThreshold",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetCompression()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEnrollmentConfigValidationError{
				field:  "Compression",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetReferenceId()) > 127 {
		return CreateEnrollmentConfigValidationError{
			field:  "ReferenceId",
			reason: "value length must be at most 127 runes",
		}
	}

	// no validation rules for NumLivenessFramesRequired

	// no validation rules for DisableServerEnrollmentTemplateStorage

	return nil
}

// CreateEnrollmentConfigValidationError is the validation error returned by
// CreateEnrollmentConfig.Validate if the designated constraints aren't met.
type CreateEnrollmentConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEnrollmentConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEnrollmentConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEnrollmentConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEnrollmentConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEnrollmentConfigValidationError) ErrorName() string {
	return "CreateEnrollmentConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEnrollmentConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEnrollmentConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEnrollmentConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEnrollmentConfigValidationError{}

// Validate checks the field values on AuthenticateConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthenticateConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsLivenessEnabled

	if _, ok := RecognitionThreshold_name[int32(m.GetLivenessThreshold())]; !ok {
		return AuthenticateConfigValidationError{
			field:  "LivenessThreshold",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetCompression()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthenticateConfigValidationError{
				field:  "Compression",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DoIncludeToken

	// no validation rules for EnrollmentToken

	switch m.AuthId.(type) {

	case *AuthenticateConfig_EnrollmentId:

		if err := m._validateUuid(m.GetEnrollmentId()); err != nil {
			return AuthenticateConfigValidationError{
				field:  "EnrollmentId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	case *AuthenticateConfig_EnrollmentGroupId:
		// no validation rules for EnrollmentGroupId

	default:
		return AuthenticateConfigValidationError{
			field:  "AuthId",
			reason: "value is required",
		}

	}

	return nil
}

func (m *AuthenticateConfig) _validateUuid(uuid string) error {
	if matched := _video_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// AuthenticateConfigValidationError is the validation error returned by
// AuthenticateConfig.Validate if the designated constraints aren't met.
type AuthenticateConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticateConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticateConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticateConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticateConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticateConfigValidationError) ErrorName() string {
	return "AuthenticateConfigValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticateConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticateConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticateConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticateConfigValidationError{}

// Validate checks the field values on ValidateRecognitionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ValidateRecognitionConfig) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetModelName()); l < 1 || l > 255 {
		return ValidateRecognitionConfigValidationError{
			field:  "ModelName",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetUserId()); l < 1 || l > 127 {
		return ValidateRecognitionConfigValidationError{
			field:  "UserId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if _, ok := RecognitionThreshold_name[int32(m.GetThreshold())]; !ok {
		return ValidateRecognitionConfigValidationError{
			field:  "Threshold",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ValidateRecognitionConfigValidationError is the validation error returned by
// ValidateRecognitionConfig.Validate if the designated constraints aren't met.
type ValidateRecognitionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateRecognitionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateRecognitionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateRecognitionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateRecognitionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateRecognitionConfigValidationError) ErrorName() string {
	return "ValidateRecognitionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateRecognitionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateRecognitionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateRecognitionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateRecognitionConfigValidationError{}
