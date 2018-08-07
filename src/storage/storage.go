package storage

import (
	"./models"
)

// IDTask struct
type IDTask struct {
	ID   string
	Task *models.Task
}

// ITaskStorage interface
type ITaskStorage interface {
	StoreTask(t models.Task) string
	ReadTask(id string) models.Task
	ReadTasks(checkFunc func(id string) bool) []IDTask

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
