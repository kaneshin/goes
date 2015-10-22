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
		elastic.SetScheme(config.GetScheme()),
		elastic.SetMaxRetries(config.MaxRetries),
		elastic.SetSniff(config.SnifferEnabled),
		elastic.SetHealthcheck(config.HealthcheckEnabled),
		elastic.SetHttpClient(config.HTTPClient),
	)
	if err != nil {
		panic(err)
	}
	return &Elasticsearch{client}
}

func (e *Elasticsearch) CreateIndexIfNotExists(index string) (*elastic.CreateIndexResult, error) {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := e.IndexExists(index).Do()
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, nil
	}
	// Create a new index.
	createIndex, err := e.CreateIndex(index).Do()
	if err != nil {
		return nil, err
	}
	return createIndex, nil
}

func (e *Elasticsearch) PutMapping(index, typ string, mapping interface{}) (*elastic.PutMappingResponse, error) {
	putMappingService := elastic.NewPutMappingService(e.Client)
	putMappingService.AllowNoIndices(true)

	switch mapping.(type) {
	case string:
		putMappingService.BodyString(mapping.(string))
	case map[string]interface{}:
		putMappingService.BodyJson(mapping.(map[string]interface{}))
	default:
		return nil, nil
	}

	return putMappingService.Do()
}
