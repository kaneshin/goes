package elasticsearch

import (
	"net/http"
	"net/url"
)

const (
	defaultURL = "http://127.0.0.1:9200"
)

type Config struct {
	url string

	// The maximum number of times that a request will be retried for failures.
	// Defaults to -1, which defers the max retry setting to the service specific
	// configuration.
	MaxRetries int

	// The HTTP client to use when sending requests.
	HTTPClient *http.Client
}

// NewConfig returns a new Config pointer
func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetURL() string {
	return c.url
}

func (c *Config) SetURL(url string) {
	c.url = url
}

func (c *Config) GetEndpoint() string {
	return c.GetURL()
}

func (c *Config) SetEndpoint(endpoint string) {
	c.SetURL(endpoint)
}

func (c *Config) IsSSL() bool {
	url, err := url.Parse(c.url)
	if err != nil {
		return false
	}
	return url.Scheme == "https"
}
