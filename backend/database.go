package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite", "./data.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS temperatures (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		value REAL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(query)
	return err
}

func saveReading(temp float64) error {
	_, err := db.Exec("INSERT INTO temperatures (value) VALUES (?)", temp)
	return err
}
