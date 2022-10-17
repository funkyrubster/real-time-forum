package sqldb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqldb/forum.db")
	if err != nil {
		panic(err.Error())
	}

	DB = db

	return DB
}

func CloseDB() {
	DB.Close()
}
