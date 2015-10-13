package elasticsearchservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElasticsearch(t *testing.T) {
	assert := assert.New(t)

	svc := New(nil)
	assert.Nil(svc)

	config := NewConfig()
	svc = New(config)
	assert.NotNil(svc)
}
