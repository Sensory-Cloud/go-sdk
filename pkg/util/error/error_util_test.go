package error_util_test

import (
	"testing"

	error_util "github.com/Sensory-Cloud/go-sdk/pkg/util/error"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func TestError(t *testing.T) {
	code := codes.Internal
	msg := "some error message"

	err := error_util.Error(code, msg)

	errStatus, ok := status.FromError(err)
	test_util.Assert(t, ok, "Status should be initialize-able from error")
	test_util.AssertEquals(t, code, errStatus.Code())
	test_util.AssertEquals(t, msg, errStatus.Message())
	test_util.AssertEquals(t, status.Code(err), code)
}

func TestErrorf(t *testing.T) {
	code := codes.AlreadyExists
	msg := "error with substitutions: 42, 21, string"

	err := error_util.Errorf(code, "error with substitutions: %d, %d, %s", 42, 21, "string")

	errStatus, ok := status.FromError(err)
	test_util.Assert(t, ok, "Status should be initialize-able from error")
	test_util.AssertEquals(t, code, errStatus.Code())
	test_util.AssertEquals(t, msg, errStatus.Message())
	test_util.AssertEquals(t, status.Code(err), code)
}
