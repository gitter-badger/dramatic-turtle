package v1

// IDTask struct
type IDTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Task struct
type Task struct {
	Name       string     `json:"name"`
	LogEntries []LogEntry `json:"logs"`
}
