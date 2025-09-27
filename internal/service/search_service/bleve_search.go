package search_service

import (
	"SmartSearch/internal/models"
	"SmartSearch/internal/repository"
	"log"

	"github.com/blevesearch/bleve/v2"
)

var Documents = []models.Document{
	{Id: "1", Title: "Go Concurrency", Content: "Go supports concurrency via goroutines and channels."},
	{Id: "2", Title: "Python AI", Content: "Python is widely used in machine learning and AI research."},
	{Id: "3", Title: "Databases", Content: "Postgres provides full-text search_service capabilities."},
}

func IndexData(index bleve.Index) {
	err1 := index.Index(Documents[0].Id, Documents[0])
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := index.Index(Documents[1].Id, Documents[1])
	if err2 != nil {
		log.Fatal(err2)
	}
	err3 := index.Index(Documents[2].Id, Documents[2])
	if err3 != nil {
		log.Fatal(err3)
	}
}

func FullTextSearch(userQuery string, indexName string) []map[string]interface{} {
	index, _ := repository.GetIndex(indexName)
	//IndexData(index)

	query := bleve.NewQueryStringQuery(userQuery)
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatal(err)
	}

	//defer func(index bleve.Index) {
	//	err := index.Close()
	//	if err != nil {
	//
	//	}
	//}(index)

	var docs []map[string]interface{}
	for _, hit := range searchResults.Hits {
		docs = append(docs, map[string]interface{}{
			"id":    hit.ID,
			"score": hit.Score,
		})
	}

	//log.Printf("found docs: %d", searchResults.Total)
	//log.Printf("found docs: %s", searchResults.Hits)
	return docs
}
