// Package to generate detailed errors with an attached stack trace
package error_util

import (
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Custom error type that provides both a stack trace and a grpc status
type DetailedError struct {
	status *status.Status
}

// Needed to be an error type
func (e DetailedError) Error() string {
	return e.status.String()
}

// grpc uses this to recover status info from errors
func (e DetailedError) GRPCStatus() *status.Status {
	return e.status
}

// Creates an error from the standard grpc status codes and attaches a stacktrace to the error
func Errorf(code codes.Code, format string, a ...interface{}) error {
	status := status.Newf(code, format, a...)

	return DetailedError{
		status: status,
	}
}

// Creates an error from the standard grpc status codes and attaches a stacktrace to the error
func Error(code codes.Code, msg string) error {
	status := status.New(code, msg)

	return DetailedError{
		status: status,
	}
}
