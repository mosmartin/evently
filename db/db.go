package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./evently.db")
	if err != nil {
		panic("🔴 Could not connect to database!")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = createTables(db)
	if err != nil {
		panic("🔴 Could not create tables!")
	}

	return db, nil
}

// createTables creates the tables in the database. should take the database connection as an argument
func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		panic("🔴 Could not create users table!")
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		panic("🔴 Could not create events table!")
	}

	return nil
}
