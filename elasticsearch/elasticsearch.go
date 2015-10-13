package elasticsearch

import (
	"gopkg.in/olivere/elastic.v2"
)

type Elasticsearch struct {
	*elastic.Client
}

// New returns a new ElasticsearchService client.
func New(config *Config) *Elasticsearch {
	client, err := elastic.NewClient(
		elastic.SetURL(config.GetURL()),
		elastic.SetMaxRetries(config.MaxRetries),
		elastic.SetHttpClient(config.HTTPClient),
	)
	if err != nil {
		return nil
	}
	return &Elasticsearch{client}
}
