package index_service

import (
	"SmartSearch/internal/repository"
	"errors"
	"fmt"
)

func DeleteIndex(indexName string) (string, error) {
	if indexName == "" {
		return "", errors.New("index_service index_name cannot be empty")
	}

	_, err := repository.DeleteIndex(indexName)
	if err != nil {
		return "", fmt.Errorf("failed to delete index %q: %w", indexName, err)
	}

	return indexName, nil
}
