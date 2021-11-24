// module for sql functions
package db

import (
	"database/sql"

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
func GetTasks() []Task {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	tasks := []Task{}
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
		tasks = append(tasks, l)
	}
	defer db.Close()
	return tasks
}

func GetLists() []List {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	lists := []List{}
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
		lists = append(lists, l)
	}
	defer db.Close()

	return lists
}

func GetTask(id string) Task {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	task := Task{}
	rows, err := db.Query("select * from task where id=" + id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.Scan(&task.Id, &task.Data, &task.Status, &task.List_id)

		if err != nil {
			panic(err)
		}

	}
	defer db.Close()

	return task
}

func GetList(id string) List {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	list := List{}
	rows, err := db.Query("select * from list where id=" + id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.Scan(&list.Id, &list.Name)

		if err != nil {
			panic(err)
		}

	}
	defer db.Close()

	return list
}
