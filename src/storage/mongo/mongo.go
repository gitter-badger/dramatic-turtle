package mongo

import (
	"context"

	"../../core"
	"../../storage"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Mongo struct
type Mongo struct {
	mongo *mongo.Client
	db    *mongo.Database

	taskStorage *TaskStorage
	logStorage  *LogStorage
}

// Connect func
func (m *Mongo) Connect() {
	client, err := mongo.Connect(
		context.Background(),
		"mongodb://localhost:27017",
		nil)
	core.CheckErr(err)

	m.db = client.Database("dramatic-turtle")
	m.taskStorage = createTaskStorage(m)
	m.logStorage = createLogStorage(m)
}

// GetTaskStorage func
func (m *Mongo) GetTaskStorage() storage.ITaskStorage {
	return m.taskStorage
}

// GetLogStorage func
func (m *Mongo) GetLogStorage() storage.ILogStorage {
	return m.logStorage
}
