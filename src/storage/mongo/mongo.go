package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"../../core"
	"../../storage"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	taskCollectionName = "logs"
)

// Mongo struct
type Mongo struct {
	mongo *mongo.Client
	db    *mongo.Database

	taskStorage *TaskStorage
}

// Connect func
func (m *Mongo) Connect(url string, database string) {
	client, err := mongo.Connect(context.Background(), url, nil)
	core.CheckErr(err)

	m.db = client.Database(database)
	m.taskStorage = createTaskStorage(m, taskCollectionName)
}

// GetTaskStorage func
func (m *Mongo) GetTaskStorage() storage.ITaskStorage {
	return m.taskStorage
}

func (m *Mongo) getID(x *mongo.InsertOneResult) string {
	c := x.InsertedID.(objectid.ObjectID)
	return c.Hex()
}
