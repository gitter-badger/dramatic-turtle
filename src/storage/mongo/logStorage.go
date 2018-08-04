package mongo

import (
	"context"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"../models"
)

// LogStorage struct
type LogStorage struct {
	mongo *Mongo
	coll  *mongo.Collection
}

func createLogStorage(m *Mongo) *LogStorage {
	return &LogStorage{mongo: m, coll: m.db.Collection("logs")}
}

// StoreLog func
func (ls *LogStorage) StoreLog(log models.Log) {
	ls.coll.InsertOne(context.Background(), log)
}

// ReadLog func
func (ls *LogStorage) ReadLog(i int) models.Log {
	var res models.Log
	ls.coll.FindOne(context.Background(),
		bson.NewDocument(bson.EC.String("id", strconv.Itoa(i)))).Decode(&res)
	return res
}
