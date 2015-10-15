package main

import (
	"flag"
	"fmt"

	"github.com/kaneshin/goes/elasticsearch"
)

var (
	accessKeyID     = flag.String("accessKeyID", "", "AWS Access key ID")
	secretAccessKey = flag.String("secretAccessKey", "", "AWS Secret Access Key")
	region          = flag.String("region", "us-east-1", "e.g.) us-east-1, ap-northeast-1")
	endpoint        = flag.String("endpoint", "http://127.0.0.1:9200", "Endpoint")
)

func main() {
	flag.Parse()

	// Configuration
	config := elasticsearch.NewConfig()
	config.SetURL(*endpoint)
	// Disable sniffer due to die for ElasticsearchService
	// Because of nothing http information of node.
	// $ curl -X GET 'https://search-aws.ap-northeast-1.es.amazonaws.com/_nodes/http'
	// {
	//     "cluster_name": "Cluster",
	//     "nodes": {
	//         "xxxxxxxxxxxxxxxxxxxxxx": {
	//             "name": "Elasticsearch",
	//             "version": "1.5.2",
	//             "build": "62ff986"
	//         }
	//     }
	// }
	config.SnifferEnabled = false

	client := elasticsearch.New(config)
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping().Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	createIndexResult, err := client.CreateIndexIfNotExists("twitter")
	if err != nil {
		panic(err)
	}
	if createIndexResult != nil {
		fmt.Printf("%+v\n", createIndexResult)
	}

	// Index a tweet (using JSON serialization)
	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	put1, err := client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet1).
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Type("tweet").
		Id("1").
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

}

// 	// Search with a term query
// 	termQuery := elastic.NewTermQuery("user", "olivere")
// 	searchResult, err := client.Search().
// 		Index("twitter").   // search in index "twitter"
// 		Query(&termQuery).  // specify the query
// 		Sort("user", true). // sort by "user" field, ascending
// 		From(0).Size(10).   // take documents 0-9
// 		Pretty(true).       // pretty print request and response JSON
// 		Do()                // execute
// 	if err != nil {
// 		// Handle error
// 		panic(err)
// 	}
//
// 	// searchResult is of type SearchResult and returns hits, suggestions,
// 	// and all kinds of other information from Elasticsearch.
// 	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
//
// 	// Each is a convenience function that iterates over hits in a search result.
// 	// It makes sure you don't need to check for nil values in the response.
// 	// However, it ignores errors in serialization. If you want full control
// 	// over iterating the hits, see below.
// 	var ttyp Tweet
// 	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
// 		if t, ok := item.(Tweet); ok {
// 			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
// 		}
// 	}
// 	// TotalHits is another convenience function that works even when something goes wrong.
// 	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())
//
// 	// Here's how you iterate through results with full control over each step.
// 	if searchResult.Hits != nil {
// 		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
//
// 		// Iterate through results
// 		for _, hit := range searchResult.Hits.Hits {
// 			// hit.Index contains the name of the index
//
// 			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
// 			var t Tweet
// 			err := json.Unmarshal(*hit.Source, &t)
// 			if err != nil {
// 				// Deserialization failed
// 			}
//
// 			// Work with tweet
// 			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
// 		}
// 	} else {
// 		// No hits
// 		fmt.Print("Found no tweets\n")
// 	}
//
// 	// Update a tweet by the update API of Elasticsearch.
// 	// We just increment the number of retweets.
// 	update, err := client.Update().Index("twitter").Type("tweet").Id("1").
// 		Script("ctx._source.retweets += num").
// 		ScriptParams(map[string]interface{}{"num": 1}).
// 		Upsert(map[string]interface{}{"retweets": 0}).
// 		Do()
// 	if err != nil {
// 		// Handle error
// 		panic(err)
// 	}
// 	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)
//
// 	// ...
//
// 	// Delete an index.
// 	deleteIndex, err := client.DeleteIndex("twitter").Do()
// 	if err != nil {
// 		// Handle error
// 		panic(err)
// 	}
// 	if !deleteIndex.Acknowledged {
// 		// Not acknowledged
// 	}
// }
