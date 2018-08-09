package v1

// GetLogEntryForTaskRequest struct
type GetLogEntryForTaskRequest struct {
}

// GetLogEntryForTaskResponse struct
type GetLogEntryForTaskResponse struct {
	Entry LogEntry `json:"entry"`
}
