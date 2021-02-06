package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Documents struct {
	Documents []Player
}

func ReadDocuments(filename string) Documents {
	lines, _ := ioutil.ReadFile(filename)
	documents := []Player{}
	for _, line := range bytes.Split(lines, []byte{'\n'}) {
		var player Player
		if err := json.Unmarshal(line, &player); err != nil {
			panic(err)
		}
		documents = append(documents, player)
	}
	return Documents{documents}
}
