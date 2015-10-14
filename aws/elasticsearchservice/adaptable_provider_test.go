package elasticsearchservice

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	assert := assert.New(t)

	provider := NewAdaptableProvider("id", "secret", "token")
	assert.NotNil(provider)
	v, err := provider.Retrieve()
	assert.Equal(credentials.Value{"id", "secret", "token"}, v)
	assert.NoError(err)
	assert.False(provider.IsExpired())
	assert.True(provider.retrieved)
}
