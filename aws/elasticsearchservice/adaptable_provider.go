package elasticsearchservice

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type AdaptableProvider struct {
	retrieved bool
	static    *credentials.StaticProvider
	env       *credentials.EnvProvider
}

// NewAdaptableProvider returns a pointer to a new AdaptableProvider object
// wrapping static and environment providers.
func NewAdaptableProvider(id, secret, token string) *AdaptableProvider {
	return &AdaptableProvider{
		retrieved: (id == "" && secret == ""),
		static: &credentials.StaticProvider{Value: credentials.Value{
			AccessKeyID:     id,
			SecretAccessKey: secret,
			SessionToken:    token,
		}},
		env: &credentials.EnvProvider{},
	}
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (p *AdaptableProvider) Retrieve() (credentials.Value, error) {
	if v, err := p.static.Retrieve(); err == nil {
		p.retrieved = true
		return v, nil
	}

	p.retrieved = false
	v, err := p.env.Retrieve()
	if err == nil {
		p.retrieved = true
		return v, nil
	}

	return credentials.Value{}, err
}

// IsExpired returns if the credentials have been retrieved.
func (p *AdaptableProvider) IsExpired() bool {
	return !p.retrieved
}
