package search_service

import (
	"SmartSearch/internal/repository"
	"log"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"
)

func FullTextSearch(userQuery string, filters map[string]interface{}, indexName string) []map[string]interface{} {
	index, _ := repository.GetIndex(indexName)

	filterQuery, _ := filterModule(filters)
	var finalQuery query.Query
	textQuery := bleve.NewQueryStringQuery(userQuery)
	if filterQuery != nil {
		finalQuery = bleve.NewConjunctionQuery(textQuery, filterQuery)
	} else {
		finalQuery = textQuery
	}

	search := bleve.NewSearchRequest(finalQuery)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		log.Println(err)
	}

	var docs []map[string]interface{}
	for _, hit := range searchResults.Hits {
		doc := map[string]interface{}{
			"id":    hit.ID,
			"score": hit.Score,
		}
		for k, v := range hit.Fields {
			doc[k] = v
		}
		docs = append(docs, doc)
	}

	return docs
}

func filterModule(filters map[string]interface{}) (*query.ConjunctionQuery, error) {
	var queries []query.Query

	for field, value := range filters {
		switch v := value.(type) {
		case string:
			q := bleve.NewMatchQuery(v)
			q.SetField(field)
			q.SetFuzziness(1)
			queries = append(queries, q)
		case float64:
			q := bleve.NewNumericRangeQuery(&v, &v)
			q.SetField(field)
			queries = append(queries, q)
		case int:
			f := float64(v)
			q := bleve.NewNumericRangeQuery(&f, &f)
			q.SetField(field)
			queries = append(queries, q)
		case int64:
			f := float64(v)
			q := bleve.NewNumericRangeQuery(&f, &f)
			q.SetField(field)
			queries = append(queries, q)
		case []float64:
			if len(v) == 2 {
				q := bleve.NewNumericRangeQuery(&v[0], &v[1])
				q.SetField(field)
				queries = append(queries, q)
			}
		}
	}

	if len(queries) == 0 {
		return nil, nil
	} else if len(queries) == 1 {
		return bleve.NewConjunctionQuery(queries[0]), nil
	}

	return bleve.NewConjunctionQuery(queries...), nil
}
