// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/event/event.proto

package event

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

	_ = common.UsageEventType(0)

	_ = common.ModelType(0)

	_ = common.UsageEventType(0)

	_ = common.ModelType(0)

	_ = common.ModelType(0)
)

// define the regex for a UUID once up-front
var _event_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PublishUsageEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PublishUsageEventsRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PublishUsageEventsRequestValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PublishUsageEventsRequestValidationError is the validation error returned by
// PublishUsageEventsRequest.Validate if the designated constraints aren't met.
type PublishUsageEventsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishUsageEventsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishUsageEventsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishUsageEventsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishUsageEventsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishUsageEventsRequestValidationError) ErrorName() string {
	return "PublishUsageEventsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PublishUsageEventsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishUsageEventsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishUsageEventsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishUsageEventsRequestValidationError{}

// Validate checks the field values on UsageEvent with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UsageEvent) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimestamp() == nil {
		return UsageEventValidationError{
			field:  "Timestamp",
			reason: "value is required",
		}
	}

	if m.GetDuration() < 0 {
		return UsageEventValidationError{
			field:  "Duration",
			reason: "value must be greater than or equal to 0",
		}
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return UsageEventValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetClientId()); l < 1 || l > 127 {
		return UsageEventValidationError{
			field:  "ClientId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if _, ok := common.UsageEventType_name[int32(m.GetType())]; !ok {
		return UsageEventValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
	}

	if l := utf8.RuneCountInString(m.GetRoute()); l < 1 || l > 511 {
		return UsageEventValidationError{
			field:  "Route",
			reason: "value length must be between 1 and 511 runes, inclusive",
		}
	}

	// no validation rules for AudioDurationMs

	// no validation rules for VideoFrameCount

	// no validation rules for TenantId

	// no validation rules for BillableFunction

	// no validation rules for TokenCount

	return nil
}

func (m *UsageEvent) _validateUuid(uuid string) error {
	if matched := _event_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UsageEventValidationError is the validation error returned by
// UsageEvent.Validate if the designated constraints aren't met.
type UsageEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageEventValidationError) ErrorName() string { return "UsageEventValidationError" }

// Error satisfies the builtin error interface
func (e UsageEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageEventValidationError{}

// Validate checks the field values on UsageEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UsageEventResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimestamp() == nil {
		return UsageEventResponseValidationError{
			field:  "Timestamp",
			reason: "value is required",
		}
	}

	if m.GetDuration() < 0 {
		return UsageEventResponseValidationError{
			field:  "Duration",
			reason: "value must be greater than or equal to 0",
		}
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return UsageEventResponseValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetClientId()); l < 1 || l > 127 {
		return UsageEventResponseValidationError{
			field:  "ClientId",
			reason: "value length must be between 1 and 127 runes, inclusive",
		}
	}

	if _, ok := common.UsageEventType_name[int32(m.GetType())]; !ok {
		return UsageEventResponseValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
	}

	if l := utf8.RuneCountInString(m.GetRoute()); l < 1 || l > 511 {
		return UsageEventResponseValidationError{
			field:  "Route",
			reason: "value length must be between 1 and 511 runes, inclusive",
		}
	}

	// no validation rules for BillableValue

	// no validation rules for BillableUnits

	// no validation rules for TenantId

	// no validation rules for BillableFunction

	// no validation rules for Credits

	return nil
}

func (m *UsageEventResponse) _validateUuid(uuid string) error {
	if matched := _event_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UsageEventResponseValidationError is the validation error returned by
// UsageEventResponse.Validate if the designated constraints aren't met.
type UsageEventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageEventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageEventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageEventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageEventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageEventResponseValidationError) ErrorName() string {
	return "UsageEventResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UsageEventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageEventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageEventResponseValidationError{}

// Validate checks the field values on UsageEventListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UsageEventListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TenantId

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UsageEventListRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAfter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UsageEventListRequestValidationError{
				field:  "After",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UsageEventListRequestValidationError{
				field:  "Before",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UsageEventListRequestValidationError is the validation error returned by
// UsageEventListRequest.Validate if the designated constraints aren't met.
type UsageEventListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageEventListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageEventListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageEventListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageEventListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageEventListRequestValidationError) ErrorName() string {
	return "UsageEventListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UsageEventListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageEventListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageEventListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageEventListRequestValidationError{}

// Validate checks the field values on UsageEventListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UsageEventListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UsageEventListResponseValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UsageEventListResponseValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UsageEventListResponseValidationError is the validation error returned by
// UsageEventListResponse.Validate if the designated constraints aren't met.
type UsageEventListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageEventListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageEventListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageEventListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageEventListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageEventListResponseValidationError) ErrorName() string {
	return "UsageEventListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UsageEventListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageEventListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageEventListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageEventListResponseValidationError{}

// Validate checks the field values on GlobalEventSummaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GlobalEventSummaryRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAfter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GlobalEventSummaryRequestValidationError{
				field:  "After",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GlobalEventSummaryRequestValidationError{
				field:  "Before",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GlobalEventSummaryRequestValidationError is the validation error returned by
// GlobalEventSummaryRequest.Validate if the designated constraints aren't met.
type GlobalEventSummaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GlobalEventSummaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GlobalEventSummaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GlobalEventSummaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GlobalEventSummaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GlobalEventSummaryRequestValidationError) ErrorName() string {
	return "GlobalEventSummaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GlobalEventSummaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGlobalEventSummaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GlobalEventSummaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GlobalEventSummaryRequestValidationError{}

// Validate checks the field values on UsageEventSummary with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UsageEventSummary) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSummaries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UsageEventSummaryValidationError{
					field:  fmt.Sprintf("Summaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UsageEventSummaryValidationError is the validation error returned by
// UsageEventSummary.Validate if the designated constraints aren't met.
type UsageEventSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageEventSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageEventSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageEventSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageEventSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageEventSummaryValidationError) ErrorName() string {
	return "UsageEventSummaryValidationError"
}

// Error satisfies the builtin error interface
func (e UsageEventSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageEventSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageEventSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageEventSummaryValidationError{}

// Validate checks the field values on UsageEventModelSummary with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UsageEventModelSummary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BillableFunction

	// no validation rules for Units

	// no validation rules for Value

	// no validation rules for Count

	// no validation rules for Credits

	// no validation rules for TenantId

	return nil
}

// UsageEventModelSummaryValidationError is the validation error returned by
// UsageEventModelSummary.Validate if the designated constraints aren't met.
type UsageEventModelSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageEventModelSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageEventModelSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageEventModelSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageEventModelSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageEventModelSummaryValidationError) ErrorName() string {
	return "UsageEventModelSummaryValidationError"
}

// Error satisfies the builtin error interface
func (e UsageEventModelSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageEventModelSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageEventModelSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageEventModelSummaryValidationError{}

// Validate checks the field values on PublishUsageEventsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PublishUsageEventsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PublishUsageEventsResponseValidationError is the validation error returned
// by PublishUsageEventsResponse.Validate if the designated constraints aren't met.
type PublishUsageEventsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishUsageEventsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishUsageEventsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishUsageEventsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishUsageEventsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishUsageEventsResponseValidationError) ErrorName() string {
	return "PublishUsageEventsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PublishUsageEventsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishUsageEventsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishUsageEventsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishUsageEventsResponseValidationError{}
