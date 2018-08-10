package v1

// GETLogTaskRequest struct
type GETLogTaskRequest struct {
}

// GETLogTaskResponse struct
type GETLogTaskResponse struct {
	Tasks []IDTask `json:"tasks"`
}
