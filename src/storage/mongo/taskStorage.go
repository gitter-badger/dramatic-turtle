package mongo

import (
	"context"

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
func (ts *TaskStorage) ReadTask(i int) models.Task {
	return models.Task{}
}
