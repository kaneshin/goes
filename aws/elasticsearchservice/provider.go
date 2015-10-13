package elasticsearchservice

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Provider struct {
	retrieved bool
	static    *credentials.StaticProvider
	env       *credentials.EnvProvider
}

func NewProvider() *Provider {
	return &Provider{
		false,
		&credentials.StaticProvider{},
		&credentials.EnvProvider{},
	}
}

func NewStaticProvider(id, secret, token string) *Provider {
	return &Provider{
		true,
		&credentials.StaticProvider{Value: credentials.Value{
			AccessKeyID:     id,
			SecretAccessKey: secret,
			SessionToken:    token,
		}},
		&credentials.EnvProvider{},
	}
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (p *Provider) Retrieve() (credentials.Value, error) {
	if v, err := p.static.Retrieve(); err == nil {
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
func (p *Provider) IsExpired() bool {
	return !p.retrieved
}
