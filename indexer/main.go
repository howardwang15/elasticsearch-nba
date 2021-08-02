package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

func IndexDocuments(documents Documents, es *elasticsearch.Client) {
	const INDEX_NAME string = "players"
	_, err := es.Indices.Delete([]string{INDEX_NAME})
	if err != nil {
		panic(err)
	}

	mappingRaw, _ := ioutil.ReadFile("./mapping.json")
	var data interface{}
	_ = json.Unmarshal(mappingRaw, &data)
	mapping, _ := json.Marshal(data)
	_, err = es.Indices.Create(
		INDEX_NAME,
		es.Indices.Create.WithBody(strings.NewReader(string(mapping))))

	for i, doc := range documents.Documents {
		encoded, _ := json.Marshal(doc)
		es.Index(INDEX_NAME,
			strings.NewReader(string(encoded)),
			es.Index.WithDocumentID(strconv.Itoa(i+1)),
		)
	}
}

func main() {
	es := CreateESClient()
	godotenv.Load()
	baseDataPath := os.Getenv("DATA_PATH")
	log.Println(es.Info())
	documents := ReadDocuments(baseDataPath)
	IndexDocuments(documents, es)
}
