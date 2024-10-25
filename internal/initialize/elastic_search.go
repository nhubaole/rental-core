package initialize

import (
	"fmt"
	"smart-rental/global"

	"github.com/elastic/go-elasticsearch/v8"
)

func InitElasticsearch() {
	url := fmt.Sprintf("http://%s:%d", global.Config.Server.Host, global.Config.ElasticSearch.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			url,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err.Error())
	}
	global.ElasticSearch = es
}