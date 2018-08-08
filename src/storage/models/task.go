package models

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Task struct
type Task struct {
	ID   objectid.ObjectID `json:"_id,omitempty"`
	Name string            `json:"name"`
	Logs []LogEntry        `json:"logs"`
}
