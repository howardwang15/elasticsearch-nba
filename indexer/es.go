package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func CreateESClient() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	log.Println("Starting Elasticsearch client...")
	log.Println(elasticsearch.Version)
	return es
}
