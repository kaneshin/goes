package elasticsearchservice

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Credentials struct {
	*credentials.Credentials
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider credentials.Provider) *Credentials {
	return &Credentials{credentials.NewCredentials(provider)}
}

// NewStaticCredentials returns a pointer to a new Credentials object
// wrapping a static credentials value provider.
func NewStaticCredentials(id, secret, token string) *Credentials {
	return NewCredentials(NewStaticProvider(id, secret, token))
}

// Getaws gets a config of aws
func (c *Credentials) Getaws() *credentials.Credentials {
	return c.Credentials
}
