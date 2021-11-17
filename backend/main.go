package main

import (
	"net/http"

	"back/db"
)

func main() {
	//index page
	http.Handle("/", http.FileServer(http.Dir("static")))

	//Get list json
	http.HandleFunc("/api/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(db.Get_list())
	})

	http.ListenAndServe(":80", nil)
}
