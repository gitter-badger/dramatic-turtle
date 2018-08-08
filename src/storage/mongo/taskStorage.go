package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/mongodb/mongo-go-driver/bson"

	"../../core"
	"../../storage"
	"../models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// TaskStorage struct
type TaskStorage struct {
	mongo *Mongo
	coll  *mongo.Collection
}

func createTaskStorage(m *Mongo, id string) *TaskStorage {
	return &TaskStorage{mongo: m, coll: m.db.Collection(id)}
}

// StoreTask func
func (ts *TaskStorage) StoreTask(task *models.Task) string {
	id, err := ts.coll.InsertOne(context.Background(), task)
	core.CheckErr(err)
	return ts.mongo.getID(id)
}

// ReadTask func
func (ts *TaskStorage) ReadTask(id string) *models.Task {
	oid, err := objectid.FromHex(id)
	core.CheckErr(err)
	var res models.Task
	ts.coll.FindOne(context.Background(),
		bson.NewDocument(bson.EC.ObjectID("_id", oid))).Decode(&res)
	return &res
}

// ReadTasks func
func (ts *TaskStorage) ReadTasks(checkFunc func(id string) bool) []*storage.IDTask {
	ctx := context.Background()
	var res = []*storage.IDTask{}
	cursor, err := ts.coll.Find(ctx, nil)
	core.CheckErr(err)
	defer cursor.Close(ctx)

	current := bson.NewDocument()
	for cursor.Next(ctx) {
		cursor.Decode(current)
		id := current.Lookup("_id").Interface().(objectid.ObjectID).Hex()
		var currentTask models.Task
		if checkFunc(id) {
			cursor.Decode(&currentTask)
			data := storage.IDTask{
				ID:   id,
				Task: &currentTask,
			}
			res = append(res, &data)
		}
	}
	return res
}

// GetLog func
func (ts *TaskStorage) GetLog(id string) storage.ILogEntryStorage {
	return createLogEntryStorage(ts.mongo, ts.coll, id)
}
