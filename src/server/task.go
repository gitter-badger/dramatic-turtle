package server

import (
	"encoding/json"
	"net/http"

	"../core"
	"../storage/models"
	"./contracts/v1"
)

func createNewTask(w http.ResponseWriter, r *http.Request) {
	var req v1.POSTLogTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)

	task := models.Task{
		Name: req.Name,
		Logs: []models.LogEntry{},
	}

	taskID := dataStorage.GetTaskStorage().StoreTask(task)

	resp := v1.POSTLogTaskResponse{
		ID:   taskID,
		Name: task.Name,
	}

	payload, err := json.Marshal(resp)
	core.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func readAllTasks(w http.ResponseWriter, r *http.Request) {
}