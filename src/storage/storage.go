package storage

import (
	"./models"
)

// IDLogEntry struct
type IDLogEntry struct {
	ID       string
	LogEntry *models.LogEntry
}

// IDTask struct
type IDTask struct {
	ID   string
	Task *models.Task
}

// ITaskStorage interface
type ITaskStorage interface {
	StoreTask(t *models.Task) string
	ReadTask(id string) *models.Task
	ReadTasks(checkFunc func(id string) bool) []*IDTask

	GetLog(id string) ILogEntryStorage
}

// ILogEntryStorage interface
type ILogEntryStorage interface {
	AppendLogEntry(e *models.LogEntry) string
	ReadLogEntries(checkFunc func(id string) bool) []*IDLogEntry
}

// IStorage interface
type IStorage interface {
	GetTaskStorage() ITaskStorage
}
