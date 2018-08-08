package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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

	taskID := dataStorage.GetTaskStorage().StoreTask(&task)

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
	tasks := dataStorage.GetTaskStorage().ReadTasks(func(id string) bool { return true })
	respTasks := []v1.IDTask{}
	for _, e := range tasks {
		respTasks = append(respTasks, v1.IDTask{
			ID:   e.ID,
			Name: e.Task.Name,
		})
	}
	resp := v1.GETLogTaskResponse{
		Tasks: respTasks,
	}

	payload, err := json.Marshal(resp)
	core.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func readTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task := dataStorage.GetTaskStorage().ReadTask(id)

	resp := v1.CreateContractTask(task)

	payload, err := json.Marshal(resp)
	core.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func createLogEntryForTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req v1.POSTTaskByIDLogRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)

	logEntry := v1.CreateModelsLogEntry(&req.Entry)
	logID := dataStorage.GetTaskStorage().GetLog(id).Append(&logEntry)

	resp := v1.POSTTaskByIDLogResponse{
		ID: logID,
	}

	payload, err := json.Marshal(resp)
	core.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
