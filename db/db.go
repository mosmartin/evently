package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", "./evently.db")
	if err != nil {
		// return nil, err
		panic("🔴 Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	return DB, nil
}
