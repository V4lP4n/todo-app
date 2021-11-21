// module for sql functions
package db

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

const db_path = "infrastructure/base.sqlite3"

type Task struct {
	Id      int    `json:"id"`
	Data    string `json:"data"`
	Status  bool   `json:"status"`
	List_id int    `json:"list_id"`
}
type List struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// return json
func Get_tasks() []byte {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	list := []Task{}
	rows, err := db.Query("select * from task")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		l := Task{}
		err := rows.Scan(&l.Id, &l.Data, &l.Status, &l.List_id)

		if err != nil {
			panic(err)
		}
		list = append(list, l)
	}
	defer db.Close()

	res, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}

	return res
}

func Get_lists() []byte {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	list := []List{}
	rows, err := db.Query("select * from list")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		l := List{}
		err := rows.Scan(&l.Id, &l.Name)

		if err != nil {
			panic(err)
		}
		list = append(list, l)
	}
	defer db.Close()

	res, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}

	return res
}
