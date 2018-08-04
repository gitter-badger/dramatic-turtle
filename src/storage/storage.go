package storage

import (
	"./models"
)

// ITaskStorage interface
type ITaskStorage interface {
	StoreTask(t models.Task)
	ReadTask(name string) models.Task
}

// ILogStorage interface
type ILogStorage interface {
	StoreLog(l models.Log)
	ReadLog(i int) models.Log
}

// IStorage interface
type IStorage interface {
	GetTaskStorage() ITaskStorage
	GetLogStorage() ILogStorage
}
