package log_util_test

import (
	"testing"

	environment "github.com/Sensory-Cloud/go-sdk/pkg/environment"
	log_util "github.com/Sensory-Cloud/go-sdk/pkg/util/log"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestGetDevLogger(t *testing.T) {
	logger, err := log_util.GetLogger()
	test_util.AssertOk(t, err)
	test_util.Assert(t, logger != nil, "a logger should be returned")
}

func TestGetProdLogger(t *testing.T) {
	environment.Set(environment.GoEnv, "prod")
	defer environment.Set(environment.GoEnv, "local")

	logger, err := log_util.GetLogger()
	test_util.AssertOk(t, err)
	test_util.Assert(t, logger != nil, "a logger should be returned")
}

func TestGetSugaredLogger(t *testing.T) {
	logger, err := log_util.GetSugaredLogger()
	test_util.AssertOk(t, err)
	test_util.Assert(t, logger != nil, "a logger should be returned")
}
