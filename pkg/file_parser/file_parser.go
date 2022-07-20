package file_parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sensory-Cloud/go-sdk/pkg/config"
)

var (
	// Strings that will evaluate to true when parsed as a boolean value
	s            = struct{}{}
	truthyValues = map[string]struct{}{"true": s, "True": s, "TRUE": s, "T": s, "t": s, "yes": s, "Yes": s, "YES": s, "y": s, "Y": s, "1": s}
)

const (
	// The key values that are used for parsing INI files
	fqdnKey           = "fullyQualifiedDomainName"
	isSecureKey       = "isSecure"
	tenantIDKey       = "tenantID"
	enrollmentTypeKey = "enrollmentType"
	credentialKey     = "credential"
	deviceIDKey       = "deviceID"
	deviceNameKey     = "deviceName"
	// The values looked for for each enrollment type
	noneKey         = "none"
	sharedSecretKey = "sharedSecret"
	jwtKey          = "jwt"
)

// A class for parsing SDK config INI files
// This is *NOT* a full featured INI parser. For example, this implementation ignores INI sections
type FileParser struct {
	contents map[string]string
}

// Creates a new INI file parser
func NewFileParser() *FileParser {
	return &FileParser{make(map[string]string)}
}

// Reads the specified file and parses out an SDKInitConfig object
func (p *FileParser) ReadFile(filename string) (*config.SDKInitConfig, error) {

	// Open the file
	contentsDat, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse to a string
	contentsStr := string(contentsDat)
	lines := strings.Split(contentsStr, "\n")

	// Read each line
	for _, line := range lines {
		p.readLine(line)
	}

	// Parse out a final config value
	return p.parseToConfig()
}

func (p *FileParser) readLine(line string) {
	// Strip any comments in the line
	line = strings.Split(line, "#")[0]

	parts := strings.Split(line, "=")
	if len(parts) != 2 {
		// Only parse lines w/ exactly one `=` sign
		return
	}

	// Trim any extra spaces and quotes and add the parsed line to the internal dict
	key := strings.Trim(strings.TrimSpace(parts[0]), "`'\"")
	value := strings.Trim(strings.TrimSpace(parts[1]), "`'\"")
	p.contents[key] = value
}

func (p *FileParser) parseToConfig() (*config.SDKInitConfig, error) {
	fqdn, present := p.contents[fqdnKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", fqdnKey)
	}
	isSecureStr, present := p.contents[isSecureKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", isSecureKey)
	}
	_, isSecure := truthyValues[isSecureStr]
	tenant, present := p.contents[tenantIDKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", tenantIDKey)
	}
	enrollmentTypeRaw, present := p.contents[enrollmentTypeKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", enrollmentTypeKey)
	}
	enrollmentType := config.None
	switch enrollmentTypeRaw {
	case noneKey:
		break
	case sharedSecretKey:
		enrollmentType = config.SharedSecret
	case jwtKey:
		enrollmentType = config.Jwt
	default:
		return nil, fmt.Errorf("unknown value for %s: %s", enrollmentTypeKey, enrollmentTypeRaw)
	}
	credential, present := p.contents[credentialKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", credentialKey)
	}
	deviceID, present := p.contents[deviceIDKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", deviceIDKey)
	}
	deviceName, present := p.contents[deviceNameKey]
	if !present {
		return nil, fmt.Errorf("missing value for configuration: %s", deviceNameKey)
	}

	return &config.SDKInitConfig{
		FullyQualifiedDomainName: fqdn,
		IsSecure:                 isSecure,
		TenantID:                 tenant,
		EnrollmentType:           enrollmentType,
		Credential:               credential,
		DeviceID:                 deviceID,
		DeviceName:               deviceName,
	}, nil
}
