package search

import (
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
)

func CreateIndex() bleve.Index {
	if _, err := os.Stat("example.bleve"); err == nil {
		index, err := bleve.Open("example.bleve")
		if err != nil {
			log.Fatalf("failed to open existing index: %v", err)
		}
		fmt.Printf("Opened existing index: %s\n", "example.bleve")
		return index
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		panic(err)
	}
	return index
}
