package mongo

import (
	"../models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// LogEntryStorage struct
type LogEntryStorage struct {
	mongo  *Mongo
	coll   *mongo.Collection
	taskID string
}

func createLogEntryStorage(m *Mongo, coll *mongo.Collection, taskID string) *LogEntryStorage {
	return &LogEntryStorage{
		mongo:  m,
		coll:   coll,
		taskID: taskID,
	}
}

// Append func
func (les *LogEntryStorage) Append(e models.LogEntry) string {
	return ""
}
