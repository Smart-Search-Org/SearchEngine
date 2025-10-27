package requests

type CreateIndexRequest struct {
	UserId    string `json:"user_id"`
	IndexName string `json:"index_name"`
}

type DeleteIndexRequest struct {
	IndexName string `json:"index_name"`
}

type GetIndexStructureRequest struct {
	UserId    string `json:"user_id"`
	IndexName string `json:"index_name"`
}
