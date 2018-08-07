package v1

// POSTLogTaskRequest struct
type POSTLogTaskRequest struct {
	Name string `json:"name"`
}

// POSTLogTaskResponse struct
type POSTLogTaskResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
