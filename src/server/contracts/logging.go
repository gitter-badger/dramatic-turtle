package contracts

// StartLoggingRequest struct
type StartLoggingRequest struct {
	Task string `json:"task"`
}

// StartLoggingResponse struct
type StartLoggingResponse struct {
	Task string `json:"task"`
	Ref  string `json:"ref"`
}

// StopLoggingRequest struct
type StopLoggingRequest struct {
	Ref string `json:"ref"`
}

// StopLoggingResponse struct
type StopLoggingResponse struct {
	Task       string `json:"task"`
	Ref        string `json:"ref"`
	LoggedTime int    `json:"logged_time"`
}
