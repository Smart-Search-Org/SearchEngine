package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

const BasePath = "./.appdata"

var (
	indexRegistry = make(map[string]bleve.Index)
	mu            sync.RWMutex
)

func LoadIndexesFromDisk() {
	files, err := os.ReadDir(BasePath)
	if err != nil {
		log.Fatalf("failed to read indexes directory: %v", err)
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		indexPath := filepath.Join(BasePath, f.Name())
		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Printf("warning: failed to open index %s: %v", f.Name(), err)
			continue
		}

		mu.Lock()
		indexRegistry[f.Name()] = index
		mu.Unlock()

		log.Printf("loaded index: %s", f.Name())
	}
}

func CreateIndex(indexName string, mapping *mapping.IndexMappingImpl) (bleve.Index, error) {
	if indexName == "" {
		return nil, errors.New("index_service indexName cannot be empty")
	}

	indexPath := filepath.Join(BasePath, indexName)

	if err := os.MkdirAll(indexPath, 0755); err != nil {
		return nil, err
	}

	index, err := bleve.New(indexPath, mapping)
	if err != nil {
		return nil, fmt.Errorf("failed to create index_service %q: %w", indexName, err)
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := indexRegistry[indexName]; exists {
		return nil, errors.New("index_service already exists")
	}

	indexRegistry[indexName] = index
	return index, nil
}

func GetIndex(indexName string) (bleve.Index, error) {
	mu.RLock()
	defer mu.RUnlock()

	idx, ok := indexRegistry[indexName]
	if !ok {
		return nil, errors.New("index_service not found")
	}
	return idx, nil
}
