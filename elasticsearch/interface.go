package elasticsearch

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

type Interface interface {
	// AddTagsRequest(*elasticsearchservice.AddTagsInput) (*request.Request, *elasticsearchservice.AddTagsOutput)
	// AddTags(*elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error)
	// CreateElasticsearchDomainRequest(*elasticsearchservice.CreateElasticsearchDomainInput) (*request.Request, *elasticsearchservice.CreateElasticsearchDomainOutput)
	// CreateElasticsearchDomain(*elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error)
	// DeleteElasticsearchDomainRequest(*elasticsearchservice.DeleteElasticsearchDomainInput) (*request.Request, *elasticsearchservice.DeleteElasticsearchDomainOutput)
	// DeleteElasticsearchDomain(*elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error)
	// DescribeElasticsearchDomainRequest(*elasticsearchservice.DescribeElasticsearchDomainInput) (*request.Request, *elasticsearchservice.DescribeElasticsearchDomainOutput)
	// DescribeElasticsearchDomain(*elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error)
	// DescribeElasticsearchDomainConfigRequest(*elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*request.Request, *elasticsearchservice.DescribeElasticsearchDomainConfigOutput)
	// DescribeElasticsearchDomainConfig(*elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error)
	// DescribeElasticsearchDomainsRequest(*elasticsearchservice.DescribeElasticsearchDomainsInput) (*request.Request, *elasticsearchservice.DescribeElasticsearchDomainsOutput)
	// DescribeElasticsearchDomains(*elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error)
	// ListDomainNamesRequest(*elasticsearchservice.ListDomainNamesInput) (*request.Request, *elasticsearchservice.ListDomainNamesOutput)
	// ListDomainNames(*elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error)
	// ListTagsRequest(*elasticsearchservice.ListTagsInput) (*request.Request, *elasticsearchservice.ListTagsOutput)
	// ListTags(*elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error)
	// RemoveTagsRequest(*elasticsearchservice.RemoveTagsInput) (*request.Request, *elasticsearchservice.RemoveTagsOutput)
	// RemoveTags(*elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error)
	// UpdateElasticsearchDomainConfigRequest(*elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*request.Request, *elasticsearchservice.UpdateElasticsearchDomainConfigOutput)
	// UpdateElasticsearchDomainConfig(*elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error)
}
