// Package to manage log formatting.
// Ideally any logged statement will use the logger from this service.
// Under the covers go.uber.org/zap is being used
package log_util

import (
	"os"

	"go.uber.org/zap"
)

// GetLogger gets the basic logger, which is build for high performance but is feature poor
func GetLogger() (*zap.Logger, error) {
	environment, ok := os.LookupEnv("GO_ENV")
	if !ok {
		environment = "prod"
	}

	if environment == "local" {
		return zap.NewDevelopment()
	} else {
		return zap.NewProduction()
	}
}

// GetSugaredLogger gets the full-featured logger
func GetSugaredLogger() (*zap.SugaredLogger, error) {
	logger, err := GetLogger()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
