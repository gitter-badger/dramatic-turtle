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
func (m *Mongo) Connect(url string, database string) {
	client, err := mongo.Connect(context.Background(), url, nil)
	core.CheckErr(err)

	m.db = client.Database(database)
	m.taskStorage = createTaskStorage(m, "tasks")
	m.logStorage = createLogStorage(m, "logs")
}

// GetTaskStorage func
func (m *Mongo) GetTaskStorage() storage.ITaskStorage {
	return m.taskStorage
}

// GetLogStorage func
func (m *Mongo) GetLogStorage() storage.ILogStorage {
	return m.logStorage
}
