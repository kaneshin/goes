package elasticsearchservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := NewConfig()
	assert.NotNil(cfg)
	assert.NotNil(cfg.Getaws())
}
