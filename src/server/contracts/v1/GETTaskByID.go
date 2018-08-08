package v1

// GETTaskByIDRequest struct
type GETTaskByIDRequest struct {
}

// GETTaskByIDResponse struct
type GETTaskByIDResponse struct {
	Tasks Task `json:"task"`
}
