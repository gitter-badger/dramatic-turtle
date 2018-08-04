package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../config"
	"../core"
	"../storage"
	"../storage/models"
	"../storage/mongo"
	"./contracts"
	"github.com/gorilla/mux"
)

var dataStorage storage.IStorage

func getTask(w http.ResponseWriter, r *http.Request) {
	var req contracts.GetTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)

	task := dataStorage.GetTaskStorage().ReadTask(req.Name)

	resp := contracts.GetTaskResponse{
		Name: task.Name}
	payload, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var req contracts.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)

	task := models.Task{
		Name: req.Name}

	dataStorage.GetTaskStorage().StoreTask(task)

	resp := contracts.CreateTaskResponse{
		Name: task.Name}
	payload, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func startLogging(w http.ResponseWriter, r *http.Request) {
	var req contracts.StartLoggingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)

	log := models.Log{
		Task:  req.Task,
		Start: 1}

	id := dataStorage.GetLogStorage().StoreLog(log)

	resp := contracts.StartLoggingResponse{
		Task: log.Task,
		Ref:  id}

	payload, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func stopLogging(w http.ResponseWriter, r *http.Request) {
	var req contracts.StopLoggingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)

	log := dataStorage.GetLogStorage().ReadLog(req.Ref)
	log.End = 1
	dataStorage.GetLogStorage().ReplaceLog(req.Ref, log)

	resp := contracts.StopLoggingResponse{
		Task:       log.Task,
		Ref:        req.Ref,
		LoggedTime: 1}

	payload, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// Start func
func Start(conf config.Config) {
	mongoConf := conf.Storage.MongoDB[0]
	m := &mongo.Mongo{}
	m.Connect(mongoConf.Connection.URL, mongoConf.Connection.Database)
	dataStorage = m

	router := mux.NewRouter()

	router.HandleFunc("/task", getTask).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/log/start", startLogging).Methods("POST")
	router.HandleFunc("/log/stop", stopLogging).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(conf.Server.Port), router))
}
