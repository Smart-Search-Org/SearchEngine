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

	docMapping := bleve.NewDocumentMapping()
	fieldMapping := bleve.NewTextFieldMapping()
	fieldMapping.Analyzer = "en"

	fieldMapping.Store = true

	docMapping.AddFieldMappingsAt("*", fieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = docMapping

	_, err := repository.CreateIndex(indexName, mapping)
	if err != nil {
		return "", fmt.Errorf("failed to save index %q: %w", indexName, err)
	}

	return indexName, nil
}
