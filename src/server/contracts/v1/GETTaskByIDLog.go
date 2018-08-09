package v1

// GETTaskByIDLogRequest struct
type GETTaskByIDLogRequest struct {
}

// GETTaskByIDLogResponse struct
type GETTaskByIDLogResponse struct {
	ID      string   `json:"id"`
	Entries []string `json:"entries"`
}
