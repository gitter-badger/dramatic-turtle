package contracts

// StartLoggingRequest struct
type StartLoggingRequest struct {
	Task string `json:"task"`
}

// StartLoggingResponse struct
type StartLoggingResponse struct {
	Task        string `json:"task"`
	ReferenceID int    `json:"reference_id"`
}

// StopLoggingRequest struct
type StopLoggingRequest struct {
	ReferenceID int `json:"reference_id"`
}

// StopLoggingResponse struct
type StopLoggingResponse struct {
	Task        string `json:"task"`
	ReferenceID int    `json:"reference_id"`
	LoggedTime  int    `json:"logged_time"`
}
