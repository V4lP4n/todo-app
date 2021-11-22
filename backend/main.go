package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/api/tasks", getTasks)
	r.HandleFunc("/api/lists", getLists)
	// r.HandleFunc("/api/tasks/{id}", getTask).Methods("GET")
	// r.HandleFunc("/api/lists/{id}", getList).Methods("GET")

	http.ListenAndServe(":80", r)
}
