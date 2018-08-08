package v1

import (
	"time"
)

// LogEntry struct
type LogEntry struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
