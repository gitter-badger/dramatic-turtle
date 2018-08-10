package v1

// POSTTaskByIDLogRequest struct
type POSTTaskByIDLogRequest struct {
	Entry LogEntry `json:"entry"`
}

// POSTTaskByIDLogResponse struct
type POSTTaskByIDLogResponse struct {
	ID string `json:"id"`
}
