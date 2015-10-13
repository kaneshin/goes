package elasticsearchservice

import (
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

type Elasticsearch struct {
	*elasticsearchservice.ElasticsearchService
}

// New returns a new ElasticsearchService client.
func New(config *Config) *Elasticsearch {
	if config != nil {
		return &Elasticsearch{elasticsearchservice.New(config.Getaws())}
	}
	return nil
}
