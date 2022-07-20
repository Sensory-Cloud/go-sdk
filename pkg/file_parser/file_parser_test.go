package file_parser_test

import (
	"testing"

	config "github.com/Sensory-Cloud/go-sdk/pkg/config"
	file_parser "github.com/Sensory-Cloud/go-sdk/pkg/file_parser"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
)

func TestBasicFile(t *testing.T) {
	expectedConfig := config.SDKInitConfig{
		FullyQualifiedDomainName: "fqdn",
		IsSecure:                 true,
		TenantID:                 "tenant",
		Credential:               "credential",
		EnrollmentType:           config.None,
		DeviceID:                 "deviceID",
		DeviceName:               "deviceName",
	}

	parser := file_parser.NewFileParser()
	config, err := parser.ReadFile("testdata/basic.ini")
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, config, &expectedConfig)
}

func TestComments(t *testing.T) {
	expectedConfig := config.SDKInitConfig{
		FullyQualifiedDomainName: "fqdn",
		IsSecure:                 true,
		TenantID:                 "tenant",
		Credential:               "credential",
		EnrollmentType:           config.SharedSecret,
		DeviceID:                 "deviceID",
		DeviceName:               "deviceName",
	}

	parser := file_parser.NewFileParser()
	config, err := parser.ReadFile("testdata/comments.ini")
	test_util.AssertOk(t, err)
	test_util.AssertEquals(t, config, &expectedConfig)
}

func TestMissingConfig(t *testing.T) {
	parser := file_parser.NewFileParser()
	config, err := parser.ReadFile("testdata/missingConfig.ini")
	test_util.Assert(t, err != nil, "Error should be thrown on missing config")
	test_util.Assert(t, config == nil, "No config should be returned on error")
}

func TestBadEnrollmentType(t *testing.T) {
	parser := file_parser.NewFileParser()
	config, err := parser.ReadFile("testdata/badEnrollmentType.ini")
	test_util.Assert(t, err != nil, "Error should be thrown on invalid enrollment type")
	test_util.Assert(t, config == nil, "No config should be returned on error")
}

func TestMissingFile(t *testing.T) {
	parser := file_parser.NewFileParser()
	config, err := parser.ReadFile("bogus")
	test_util.Assert(t, err != nil, "Error should be thrown on missing file")
	test_util.Assert(t, config == nil, "No config should be returned on error")
}
