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
	Id     int    `json:"id"`
	Data   string `json:"data"`
	Status bool   `json:"status"`
	ListId int    `json:"list_id"`
}

func (T *Task) Build() {
	db, err := sql.Open(db_provider, db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select data, status, list_id from task where id=" + strconv.Itoa(T.Id))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&T.Data, &T.Status, &T.ListId)
		if err != nil {
			panic(err)
		}

	}

}

type TaskList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
	Tasks  []Task `json:"tasks"`
}

func (TL *TaskList) Build() {
	//Open db connection and prepare to close it
	db, err := sql.Open(db_provider, db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select name, user_id from task_list where id=" + strconv.Itoa(TL.Id))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&TL.Name, &TL.UserId)
		if err != nil {
			panic(err)
		}

	}

	//getting tasks for  tasklist
	rows, err = db.Query("select id, data, status from task where list_id=" + strconv.Itoa(TL.Id))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		T := Task{ListId: TL.Id}
		err := rows.Scan(&T.Id, &T.Data, &T.Status)
		if err != nil {
			panic(err)
		}
		TL.Tasks = append(TL.Tasks, T)

	}

}

type User struct {
	Id        int        `json:"id"`
	Uname     string     `json:"uname"`
	TaskLists []TaskList `json:"task_lists"`
}

func (U *User) Build() {
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

	//getting tasks for every tasklist
	for i := range U.TaskLists {
		rows, err = db.Query("select id, data, status from task where list_id=" + strconv.Itoa(U.TaskLists[i].Id))
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			T := Task{ListId: U.TaskLists[i].Id}
			err := rows.Scan(&T.Id, &T.Data, &T.Status)
			if err != nil {
				panic(err)
			}
			U.TaskLists[i].Tasks = append(U.TaskLists[i].Tasks, T)

		}

	}
}

type UserList struct {
	Users []User `json:"users"`
}

func (UL *UserList) Build() {
	//Open db connection and prepare to close it
	db, err := sql.Open(db_provider, db_path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Push uname from db
	rows, err := db.Query("select id from user ")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		U := User{}
		err := rows.Scan(&U.Id)
		if err != nil {
			panic(err)
		}
		U.Build()
		UL.Users = append(UL.Users, U)
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
		err := rows.Scan(&l.Id, &l.Data, &l.Status, &l.ListId)

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
