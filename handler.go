package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var tasks = make(map[string]Task)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = uuid.NewString()
	tasks[task.ID] = task
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = id
	tasks[id] = task
	json.NewEncoder(w).Encode(task)
}

func PatchTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, exists := tasks[id]
	if !exists {
		http.NotFound(w, r)
		return
	}
	var patchData map[string]string
	json.NewDecoder(r.Body).Decode(&patchData)

	if title, ok := patchData["title"]; ok {
		task.Title = title
	}
	if details, ok := patchData["details"]; ok {
		task.Details = details
	}
	tasks[id] = task
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
}
