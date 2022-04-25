// Package to generate detailed errors with an attached stack trace
package error_util

import (
	errors "github.com/go-errors/errors"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Custom error type that provides both a stack trace and a grpc status
type DetailedError struct {
	status *status.Status
	frame  []errors.StackFrame
}

// Needed to be an error type
func (e DetailedError) Error() string {
	return e.status.String()
}

// grpc uses this to recover status info from errors
func (e DetailedError) GRPCStatus() *status.Status {
	return e.status
}

// Sentry uses this to generate a stack trace
func (e DetailedError) StackFrames() []errors.StackFrame {
	return e.frame
}

// Creates an error from the standard grpc status codes and attaches a stacktrace to the error
func Errorf(code codes.Code, format string, a ...interface{}) error {
	err := status.Errorf(code, format, a...)
	wrapped := errors.Wrap(err, 1)
	status := status.Newf(code, format, a...)

	return DetailedError{
		status: status,
		frame:  wrapped.StackFrames(),
	}
}

// Creates an error from the standard grpc status codes and attaches a stacktrace to the error
func Error(code codes.Code, msg string) error {
	err := status.Error(code, msg)
	wrapped := errors.Wrap(err, 1)
	status := status.New(code, msg)

	return DetailedError{
		status: status,
		frame:  wrapped.StackFrames(),
	}
}
