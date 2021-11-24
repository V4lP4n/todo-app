package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/tasks/", GetTasks)
	r.HandleFunc("/api/lists/", GetLists)
	r.HandleFunc("/api/tasks/{id}", GetTask).Methods("GET")
	r.HandleFunc("/api/lists/{id}", GetList).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":80", r))
}
