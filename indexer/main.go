package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

func IndexDocuments(documents Documents, es *elasticsearch.Client) {
	// Get date for index creation
	now := time.Now().Format("20060102")
	indexName := "players-" + now

	// Delete index if it already exists
	_, err := es.Indices.Delete([]string{indexName})
	if err != nil {
		panic(err)
	}

	// Read mapping and create index
	mappingRaw, _ := ioutil.ReadFile("./mapping.json")
	var data interface{}
	_ = json.Unmarshal(mappingRaw, &data)
	mapping, _ := json.Marshal(data)
	_, err = es.Indices.Create(
		indexName,
		es.Indices.Create.WithBody(strings.NewReader(string(mapping))))
	if err != nil {
		panic(err)
	}

	// Index all documents
	for i, doc := range documents.Documents {
		encoded, _ := json.Marshal(doc)
		es.Index(indexName,
			strings.NewReader(string(encoded)),
			es.Index.WithDocumentID(strconv.Itoa(i+1)),
		)
	}
}

func main() {
	godotenv.Load()
	dataPath := os.Getenv("DATA_PATH")
	esHost := os.Getenv("ELASTIC_HOST")
	es := CreateESClient(esHost)
	log.Println("Reading from: " + dataPath)
	documents := ReadDocuments(dataPath)
	IndexDocuments(documents, es)
}
