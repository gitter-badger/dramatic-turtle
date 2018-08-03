package mongo

import (
	"../models"
)

// LogStorage struct
type LogStorage struct {
	mongo *Mongo
}

func createLogStorage(m *Mongo) *LogStorage {
	return &LogStorage{mongo: m}
}

// StoreLog func
func (ls *LogStorage) StoreLog(log models.Log) {

}

// ReadLog func
func (ls *LogStorage) ReadLog(i int) models.Log {
	return models.Log{}
}
