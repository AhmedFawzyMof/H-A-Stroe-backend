package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./HAstore.db")

	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
