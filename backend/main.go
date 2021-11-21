package main

import (
	"net/http"

	"back/db"
)

func main() {
	//index page
	http.Handle("/", http.FileServer(http.Dir("static")))

	//Get list json
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write(db.Get_tasks())
	})

	http.HandleFunc("/api/lists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write(db.Get_lists())
	})

	http.ListenAndServe(":80", nil)
}
