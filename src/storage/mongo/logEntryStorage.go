package mongo

import (
	"context"

	"../../core"
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
func (les *LogEntryStorage) Append(e models.LogEntry) string {
	taskID, err := objectid.FromHex(les.taskID)
	core.CheckErr(err)

	e.ID = objectid.New()
	filter := bson.NewDocument(
		bson.EC.ObjectID("_id", taskID),
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
