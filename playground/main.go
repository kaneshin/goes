package main

import (
	"flag"
	"fmt"

	goesess "github.com/kaneshin/goes/aws/elasticsearchservice"
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
	creds := goesess.NewAdaptableCredentials(
		*accessKeyID, *secretAccessKey, "")
	config := goesess.NewConfig()
	config.SetRegion(*region)
	config.SetCredentials(creds)
	esconfig := elasticsearch.NewConfig()
	esconfig.SetURL(*endpoint)
	config.Set(esconfig)

	es := goesess.New(config)
	fmt.Println(es)
}
