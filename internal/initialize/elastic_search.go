package initialize

import (
	"context"
	"fmt"
	"smart-rental/global"
	"smart-rental/internal/constants"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
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

func ESCreateIndexIfNotExist(){
	_, err :=esapi.IndicesExistsRequest{
		Index: []string{constants.SEARCH_INDEX},
	}.Do(context.Background(), global.ElasticSearch )

	if err != nil {
		global.ElasticSearch.Indices.Create(constants.SEARCH_INDEX)
	}
}