package v1

// GETLogTaskRequest struct
type GETLogTaskRequest struct {
}

// GETLogTaskResponse struct
type GETLogTaskResponse struct {
	Tasks []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"tasks"`
}
