package elasticsearchservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/kaneshin/goes/elasticsearch"
)

type Config struct {
	*aws.Config
}

// NewConfig returns a new Config pointer
func NewConfig() *Config {
	return &Config{aws.NewConfig()}
}

func (c *Config) Set(config *elasticsearch.Config) {
	c.WithDisableSSL(!config.IsSSL()).WithMaxRetries(config.MaxRetries)
	if config.HTTPClient != nil {
		c.WithHTTPClient(config.HTTPClient)
	}
	if endpoint := config.GetEndpoint(); len(endpoint) > 0 {
		c.SetEndpoint(endpoint)
	}
}

// SetCredentials sets a config Credentials value
func (c *Config) SetCredentials(creds *Credentials) {
	c.WithCredentials(creds.Getaws())
}

// SetEndpoint sets a config Endpoint value
func (c *Config) SetEndpoint(endpoint string) {
	c.WithEndpoint(endpoint)
}

// SetRegion sets a config Region value
func (c *Config) SetRegion(region string) {
	c.WithRegion(region)
}

// Getaws gets a config of aws
func (c *Config) Getaws() *aws.Config {
	return c.Config
}
