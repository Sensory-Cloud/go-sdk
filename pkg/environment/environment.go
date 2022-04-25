// Package to manage environmental configuration retrieval.
// This package will also look for a .env file and load the
// corresponding properties into the environment if found.
package environment

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Environment string

const (
	LocalEnv   Environment = "local"
	DevelopEnv Environment = "develop"
	StageEnv   Environment = "stage"
	ProdEnv    Environment = "production"
)

type EnvironmentVariable string

// #nosec G101 -- Potential hardcoded credentials
const (
	GoEnv            EnvironmentVariable = "GO_ENV"             // Golang environment
	AppToken         EnvironmentVariable = "APP_TOKEN"          // Token used to obtain OAuth credentials
	OauthTokenIssuer EnvironmentVariable = "OAUTH_TOKEN_ISSUER" // Allowed token issuer for OAuth validation
	CloudHost        EnvironmentVariable = "CLOUD_HOST"         // Sensory Cloud hostname with port prefix
)

func init() {
	envFilePath := Get("ENV_FILE_PATH", ".env")

	// Load .env file into environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		fmt.Println(err)
	}
}

// Get will return an environment variable if it is found.
// Otherwise it will return the specified fallback.
func Get(key EnvironmentVariable, fallback string) string {
	if value, ok := os.LookupEnv(string(key)); ok && value != "" {
		return strings.TrimSpace(value)
	}

	return fallback
}

// GetEnvironment will return the current running environment
func GetEnvironment() Environment {
	return Environment(strings.ToLower(Get("GO_ENV", string(ProdEnv))))
}

// Set will set an environment variable
func Set(key EnvironmentVariable, value string) error {
	return os.Setenv(string(key), value)
}

// IsTruthy will check if an environment variable is a truthy value
// For a True it accepts: 1, t, T, TRUE, true, True for True
// For a False it accepts: 0, f, F, FALSE, false, False.
func IsTruthy(key EnvironmentVariable, fallback bool) bool {
	truthy, _ := strconv.ParseBool(Get(key, strconv.FormatBool(fallback)))
	return truthy
}
