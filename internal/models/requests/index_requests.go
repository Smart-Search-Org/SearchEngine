package requests

type CreateIndexRequest struct {
	IndexName string `json:"index_name"`
}

type DeleteIndexRequest struct {
	IndexName string `json:"index_name"`
}

