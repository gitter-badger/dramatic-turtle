package server

import (
	"log"
	"net/http"
	"strconv"

	"../config"
	"../storage"
	"../storage/mongo"
	"github.com/gorilla/mux"
)

var dataStorage storage.IStorage

func stub(w http.ResponseWriter, r *http.Request) {
}

// Start func
func Start(conf config.Config) {
	mongoConf := conf.Storage.MongoDB
	m := &mongo.Mongo{}
	m.Connect(mongoConf.Connection.URL, mongoConf.Connection.Database)
	dataStorage = m

	router := mux.NewRouter()

	router.HandleFunc("/v1/log/task", createNewTask).Methods("POST")
	router.HandleFunc("/v1/log/task", readAllTasks).Methods("GET")
	//router.HandleFunc("/v1/log/task/{id}", stub).Methods("POST")
	//router.HandleFunc("/v1/log/task/{id}", stub).Methods("GET")

	//router.HandleFunc("/v1/log/task/{id}/entry", stub).Methods("POST")
	//router.HandleFunc("/v1/log/task/{id}/entry", stub).Methods("GET")
	//router.HandleFunc("/v1/log/task/{id}/entry/{entryID}", stub).Methods("POST")
	//router.HandleFunc("/v1/log/task/{id}/entry/{entryID}", stub).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(conf.Server.Port), router))
}
