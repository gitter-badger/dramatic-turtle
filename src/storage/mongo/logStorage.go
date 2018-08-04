package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"../../core"
	"../models"
)

// LogStorage struct
type LogStorage struct {
	mongo *Mongo
	coll  *mongo.Collection
}

func createLogStorage(m *Mongo, n string) *LogStorage {
	return &LogStorage{mongo: m, coll: m.db.Collection(n)}
}

// StoreLog func
func (ls *LogStorage) StoreLog(log models.Log) string {
	id, err := ls.coll.InsertOne(context.Background(), log)
	core.CheckErr(err)
	return ls.mongo.getID(id)
}

// ReadLog func
func (ls *LogStorage) ReadLog(id string) models.Log {
	var res models.Log
	x, err := objectid.FromHex(id)
	core.CheckErr(err)
	ls.coll.FindOne(context.Background(),
		bson.NewDocument(bson.EC.ObjectID("_id", x))).Decode(&res)
	return res
}

// ReplaceLog func
func (ls *LogStorage) ReplaceLog(id string, log models.Log) {
	oid, err := objectid.FromHex(id)
	core.CheckErr(err)
	_, err = ls.coll.ReplaceOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", oid),
		), log)
	core.CheckErr(err)
}
