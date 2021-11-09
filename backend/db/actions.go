// module for sql functions
package db

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

const db_path = "infrastructure/base.sqlite3"

type Todo struct {
	Id     int    `json:"id"`
	Data   string `json:"data"`
	Status bool   `json:"status"`
}

// return json
func Get_list() []byte {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}
	list := []Todo{}
	rows, err := db.Query("select * from todo")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		l := Todo{}
		err := rows.Scan(&l.Id, &l.Data, &l.Status)

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
