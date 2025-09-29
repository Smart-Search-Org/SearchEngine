package requests

type SearchRequest struct {
	IndexName string                 `json:"index_name"`
	Query     string                 `json:"query"`
	Filters   map[string]interface{} `json:"filters"`
}
