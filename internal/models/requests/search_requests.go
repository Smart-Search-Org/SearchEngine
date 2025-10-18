package requests

type SearchRequest struct {
	UserId    string                 `json:"user_id"`
	IndexName string                 `json:"index_name"`
	Query     string                 `json:"query"`
	Filters   map[string]interface{} `json:"filters"`
}
