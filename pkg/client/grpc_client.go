// Package useful for common client-side Getters
package grpc_client

import (
	"context"
	"crypto/tls"
	"strings"
	"time"

	log_util "github.com/Sensory-Cloud/go-sdk/pkg/util/log"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	log, _ = log_util.GetSugaredLogger()
)

const (
	DefaultTimeout = 8000 * time.Second
)

// GetDefaultContext returns the default context used for most short-lived client-side requests
func GetDefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 8000*time.Second)
}

// NewCloudClient returns a connection to a grpc client at host specified by the cloudHost input string.
// A cleanup function is returned. This function should be called once your code is done with the client connection.
func NewCloudClient(cloudHost string, opts ...grpc.DialOption) (*grpc.ClientConn, func(), error) {
	if isInsecureHost(cloudHost) {
		opts = append(opts, grpc.WithInsecure())
	} else {
		// TODO: Improve TLS security requirements to prevent machine-in-the-middle attacks
		// https://github.com/denji/golang-tls
		// #nosec G402 -- https://sensory.atlassian.net/browse/VS2-375
		config := &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	}

	log.Infof("Connecting to %s", cloudHost)
	client, err := grpc.Dial(cloudHost, opts...)
	if err != nil {
		return nil, func() {}, err
	}

	cleanup := func() {
		err := client.Close()
		if err != nil {
			log.Errorf("Failed to close grpc client: %s", err.Error())
		}
	}

	return client, cleanup, nil
}

func isInsecureHost(host string) bool {
	return strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "localhost") || strings.HasPrefix(host, "titan")
}
