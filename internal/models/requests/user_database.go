package requests

type PopulateIndexRequest struct {
	UserId    string `json:"user_id"`
	Driver    string `json:"driver"`
	DSN       string `json:"dsn"`
	Table     string `json:"table"`
	IndexName string `json:"index_name"`
}
