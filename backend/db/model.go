// module for sql functions
package db

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const db_path = "infrastructure/base.sqlite3"
const db_provider = "sqlite3"

type Task struct {
	Id      int    `json:"id"`
	Data    string `json:"data"`
	Status  bool   `json:"status"`
	List_id int    `json:"list_id"`
}
type TaskList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
	Tasks  []Task `json:"tasks"`
}

type User struct {
	Id        int        `json:"id"`
	Uname     string     `json:"uname"`
	TaskLists []TaskList `json:"task_lists"`
}

func (U *User) BuildUser() {
	//Open db connection and prepare to close it
	db, err := sql.Open(db_provider, db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Push uname from db
	rows, err := db.Query("select uname from user where id=" + strconv.Itoa(U.Id))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&U.Uname)
		if err != nil {
			panic(err)
		}
	}
	//getting User task list

	rows, err = db.Query("select id, name from task_list where user_id=" + strconv.Itoa(U.Id))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		TL := TaskList{UserId: U.Id}
		err := rows.Scan(&TL.Id, &TL.Name)
		if err != nil {
			panic(err)
		}
		U.TaskLists = append(U.TaskLists, TL)
	}

	for i := range U.TaskLists {
		rows, err = db.Query("select id, data, status from task where list_id=" + strconv.Itoa(U.TaskLists[i].Id))
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			T := Task{List_id: U.TaskLists[i].Id}
			err := rows.Scan(&T.Id, &T.Data, &T.Status)
			if err != nil {
				panic(err)
			}
			U.TaskLists[i].Tasks = append(U.TaskLists[i].Tasks, T)

		}

	}
}

func GetTasks() []Task {
	db, err := sql.Open(db_provider, db_path)
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

func GetLists() []TaskList {
	db, err := sql.Open(db_provider, db_path)
	if err != nil {
		panic(err)
	}
	lists := []TaskList{}
	rows, err := db.Query("select * from task_list")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		l := TaskList{}
		err := rows.Scan(&l.Id, &l.Name, &l.UserId)

		if err != nil {
			panic(err)
		}
		lists = append(lists, l)
	}
	defer db.Close()

	return lists
}

func GetTask(id string) Task {
	db, err := sql.Open(db_provider, db_path)
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

func GetList(id string) TaskList {
	db, err := sql.Open(db_provider, db_path)
	if err != nil {
		panic(err)
	}
	list := TaskList{}
	rows, err := db.Query("select id, name, user_id from task_list where id=" + id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.Scan(&list.Id, &list.Name, &list.UserId)

		if err != nil {
			panic(err)
		}

	}
	defer db.Close()

	return list
}
