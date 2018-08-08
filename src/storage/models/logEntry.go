package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// LogEntry struct
type LogEntry struct {
	ID    objectid.ObjectID `json:"_id,omitempty"`
	Start time.Time         `json:"start"`
	End   time.Time         `json:"end"`
}
