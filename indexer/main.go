package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

func IndexDocuments(documents Documents, es *elasticsearch.Client) {
	for i, doc := range documents.Documents {
		encoded, _ := json.Marshal(doc)
		req := esapi.IndexRequest{
			Index:      "test",
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(string(encoded)),
		}
		_, err := req.Do(context.Background(), es)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	es := CreateESClient()
	log.Println(es.Info())
	documents := ReadDocuments("../data/players.json")
	IndexDocuments(documents, es)
}
