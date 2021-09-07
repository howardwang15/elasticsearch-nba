package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
)

func CreateESClient(esHost string) *elasticsearch.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	config := elasticsearch.Config{
		Addresses: []string{esHost + ":9200"},
		Username:  os.Getenv("ELASTIC_USERNAME"),
		Password:  os.Getenv("ELASTIC_PASSWORD"),
		Logger:    &estransport.ColorLogger{Output: os.Stdout},
		Transport: transport,
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	log.Println("Starting Elasticsearch client...version", elasticsearch.Version)
	log.Println(es.Cluster.Health())
	return es
}
