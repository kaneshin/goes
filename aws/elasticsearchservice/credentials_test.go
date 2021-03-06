package elasticsearchservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentials(t *testing.T) {
	assert := assert.New(t)

	creds := NewAdaptableCredentials("", "", "")
	assert.NotNil(creds)
	assert.NotNil(creds.Getaws())
}
