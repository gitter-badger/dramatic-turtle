package contracts

// CreateTaskRequest struct
type CreateTaskRequest struct {
	Name string `json:"name"`
}

// CreateTaskResponse struct
type CreateTaskResponse struct {
	Name string `json:"name"`
}

// GetTaskRequest struct
type GetTaskRequest struct {
	Name string `json:"name"`
}

// GetTaskResponse struct
type GetTaskResponse struct {
	Name string `json:"name"`
}
