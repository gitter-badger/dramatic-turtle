package mongo

import (
	"context"

	"../../core"
	"../../storage"
	"../models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
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
func (les *LogEntryStorage) AppendLogEntry(e *models.LogEntry) string {
	taskID, err := objectid.FromHex(les.taskID)
	core.CheckErr(err)

	e.ID = objectid.New()
	filter := bson.NewDocument(
		bson.EC.ObjectID("id", taskID),
	)
	update := bson.NewDocument(
		bson.EC.SubDocumentFromElements("$push",
			bson.EC.Interface("logs", e),
		),
	)
	_, err = les.coll.UpdateOne(context.Background(), filter, update)
	core.CheckErr(err)

	return e.ID.Hex()
}

// ReadLogEntry func
func (les *LogEntryStorage) ReadLogEntry(id string) *models.LogEntry {
	task := les.mongo.GetTaskStorage().ReadTask(les.taskID)

	for _, le := range task.Logs {
		if id == le.ID.Hex() {
			return &le
		}
	}
	return &models.LogEntry{}
}

// ReadLogEntries func
func (les *LogEntryStorage) ReadLogEntries(checkFunc func(id string) bool) []*storage.IDLogEntry {
	task := les.mongo.GetTaskStorage().ReadTask(les.taskID)

	entries := []*storage.IDLogEntry{}
	for _, le := range task.Logs {
		if checkFunc(le.ID.Hex()) {
			entries = append(entries, &storage.IDLogEntry{
				ID:       le.ID.Hex(),
				LogEntry: &le,
			})
		}
	}
	return entries
}
