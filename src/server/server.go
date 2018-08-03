package server

import (
	"encoding/json"
	"fmt"
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

var storages []storage.IStorage

func getTask(w http.ResponseWriter, r *http.Request) {
	var data contracts.GetTaskRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	core.CheckErr(err)

	fmt.Println(data.Name)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var data contracts.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	core.CheckErr(err)
	fmt.Println(data.Name)

	task := models.Task{
		Name: data.Name}

	for _, s := range storages {
		s.GetTaskStorage().StoreTask(task)
	}
}

func startLogging(w http.ResponseWriter, r *http.Request) {
	var data contracts.StartLoggingRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	core.CheckErr(err)

	fmt.Println(data.Task)
}

func stopLogging(w http.ResponseWriter, r *http.Request) {
	var data contracts.StopLoggingRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	core.CheckErr(err)

	fmt.Println(strconv.Itoa(data.ReferenceID))
}

// Start func
func Start(conf config.Config) {
	m := &mongo.Mongo{}
	m.Connect()
	storages = append(storages, m)

	router := mux.NewRouter()

	router.HandleFunc("/task", getTask).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/log/start", startLogging).Methods("POST")
	router.HandleFunc("/log/stop", stopLogging).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(conf.Server.Port), router))
}
