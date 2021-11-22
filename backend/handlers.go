package main

import (
	"back/db"
	"encoding/json"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data := db.Get_tasks()
	json.NewEncoder(w).Encode(data)
}
func getLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := db.Get_lists()
	json.NewEncoder(w).Encode(data)

}
func getTask(w http.ResponseWriter, r *http.Request) {

}
func getList(w http.ResponseWriter, r *http.Request) {

}
