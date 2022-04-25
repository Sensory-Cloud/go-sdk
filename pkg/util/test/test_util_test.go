package test_util_test

import (
	"fmt"
	"testing"
	"time"

	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAssert(t *testing.T) {
	mockTest := &testing.T{}
	go test_util.Assert(mockTest, false, "this test should fail")
	testFailed := pollForTestFailure(mockTest)

	if !testFailed {
		fmt.Println("test was expected to fail")
		t.FailNow()
	}

	mockTest = &testing.T{}
	go test_util.Assert(mockTest, true, "this test should fail")
	testFailed = pollForTestFailure(mockTest)

	if testFailed {
		fmt.Println("test was expected to pass")
		t.FailNow()
	}
}

func TestAssertOk(t *testing.T) {
	mockTest := &testing.T{}
	go test_util.AssertOk(mockTest, status.Error(codes.Internal, "error!"))
	testFailed := pollForTestFailure(mockTest)

	if !testFailed {
		fmt.Println("test was expected to fail")
		t.FailNow()
	}

	mockTest = &testing.T{}
	go test_util.AssertOk(mockTest, nil)
	testFailed = pollForTestFailure(mockTest)

	if testFailed {
		fmt.Println("test was expected to pass")
		t.FailNow()
	}
}

func TestAssertEquals(t *testing.T) {
	mockTest := &testing.T{}
	go test_util.AssertEquals(mockTest, 1, 1)
	testFailed := pollForTestFailure(mockTest)

	if testFailed {
		fmt.Println("test was expected to pass")
		t.FailNow()
	}

	mockTest = &testing.T{}
	go test_util.AssertEquals(mockTest, 1, 2)
	testFailed = pollForTestFailure(mockTest)

	if !testFailed {
		fmt.Println("test was expected to fail")
		t.FailNow()
	}

	mockTest = &testing.T{}
	go test_util.AssertEquals(mockTest, "good", "good")
	testFailed = pollForTestFailure(mockTest)

	if testFailed {
		fmt.Println("test was expected to pass")
		t.FailNow()
	}

	mockTest = &testing.T{}
	go test_util.AssertEquals(mockTest, "good", "bad")
	testFailed = pollForTestFailure(mockTest)

	if !testFailed {
		fmt.Println("test was expected to fail")
		t.FailNow()
	}
}

func pollForTestFailure(mockTest *testing.T) bool {
	tickerCheck := time.NewTicker(5 * time.Millisecond)
	tickerStop := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-tickerCheck.C:
			if mockTest.Failed() {
				return true
			}
		case <-tickerStop.C:
			tickerCheck.Stop()
			tickerStop.Stop()
			return mockTest.Failed()
		}
	}
}
