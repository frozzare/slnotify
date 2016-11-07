package httputil

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// Client represents the custom http client.
type Client struct {
	*http.Client
}

// Config represents client config.
type Config struct {
	DualStack  bool
	SkipVerify bool
	Timeout    time.Duration
}

// NewClient will create a new http client instance.
func NewClient(config *Config) *Client {
	if config == nil {
		config = &Config{}
	}

	t := &http.Transport{
		Dial: (&net.Dialer{
			DualStack: config.DualStack,
			Timeout:   config.Timeout,
		}).Dial,
	}

	if config.SkipVerify {
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return &Client{&http.Client{Transport: t}}
}
