package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"

	"../models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// TaskStorage struct
type TaskStorage struct {
	mongo *Mongo
	coll  *mongo.Collection
}

func createTaskStorage(m *Mongo) *TaskStorage {
	return &TaskStorage{mongo: m, coll: m.db.Collection("tasks")}
}

// StoreTask func
func (ts *TaskStorage) StoreTask(task models.Task) {
	ts.coll.InsertOne(context.Background(), task)
}

// ReadTask func
func (ts *TaskStorage) ReadTask(name string) models.Task {
	var res models.Task
	ts.coll.FindOne(context.Background(),
		bson.NewDocument(bson.EC.String("name", name))).Decode(&res)
	return res
}
