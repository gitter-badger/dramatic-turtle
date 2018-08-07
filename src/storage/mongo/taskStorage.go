package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"

	"../../core"
	"../../storage"
	"../models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// TaskStorage struct
type TaskStorage struct {
	mongo *Mongo
	coll  *mongo.Collection
}

func createTaskStorage(m *Mongo, id string) *TaskStorage {
	return &TaskStorage{mongo: m, coll: m.db.Collection(id)}
}

// StoreTask func
func (ts *TaskStorage) StoreTask(task models.Task) string {
	id, err := ts.coll.InsertOne(context.Background(), task)
	core.CheckErr(err)
	return ts.mongo.getID(id)
}

// ReadTask func
func (ts *TaskStorage) ReadTask(id string) models.Task {
	var res models.Task
	ts.coll.FindOne(context.Background(),
		bson.NewDocument(bson.EC.String("name", id))).Decode(&res)
	return res
}

// GetLog func
func (ts *TaskStorage) GetLog(id string) storage.ILogEntryStorage {
	return createLogEntryStorage(ts.mongo, ts.coll, id)
}
