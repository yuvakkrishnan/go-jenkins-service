package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", PatchTask).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	log.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
