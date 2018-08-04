package storage

import (
	"./models"
)

// ITaskStorage interface
type ITaskStorage interface {
	StoreTask(t models.Task) string
	ReadTask(name string) models.Task
}

// ILogStorage interface
type ILogStorage interface {
	StoreLog(l models.Log) string
	ReplaceLog(id string, l models.Log)
	ReadLog(id string) models.Log
}

// IStorage interface
type IStorage interface {
	GetTaskStorage() ITaskStorage
	GetLogStorage() ILogStorage
}
