package storage

import (
	"./models"
)

// ITaskStorage interface
type ITaskStorage interface {
	StoreTask(t models.Task) string
	ReadTask(id string) models.Task

	GetLog(id string) ILogEntryStorage
}

// ILogEntryStorage interface
type ILogEntryStorage interface {
	Append(e models.LogEntry) string
}

// IStorage interface
type IStorage interface {
	GetTaskStorage() ITaskStorage
}
