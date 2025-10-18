package index_doc_service

import (
	"SmartSearch/internal/repository/index"
	"fmt"
	"strconv"
)

func AddDocs(indexName string, raw []map[string]interface{}) error {
	index, err := index.GetIndex(indexName)
	if err != nil {
		return err
	}

	batch := index.NewBatch()

	for i, doc := range raw {
		var id string
		switch v := doc["id"].(type) {
		case string:
			id = v
		case int:
			id = strconv.Itoa(v)
		case int64:
			id = strconv.FormatInt(v, 10)
		case float64:
			id = strconv.Itoa(int(v))
		default:
			id = fmt.Sprintf("doc-%d", i)
		}

		if err := batch.Index(id, doc); err != nil {
			return err
		}
	}

	return index.Batch(batch)
}
