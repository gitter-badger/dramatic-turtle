package models

// Task struct
type Task struct {
	Name string     `json:"name"`
	Logs []LogEntry `json:"logs"`
}
