package index_service

import (
	"SmartSearch/internal/repository"
	"errors"
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

func CreateIndex(indexName string) (string, error) {
	if indexName == "" {
		return "", errors.New("index_service index_name cannot be empty")
	}

	mapping := bleve.NewIndexMapping()

	_, err := repository.CreateIndex(indexName, mapping)
	if err != nil {
		return "", fmt.Errorf("failed to save index_service %q: %w", indexName, err)
	}

	return indexName, nil
}
