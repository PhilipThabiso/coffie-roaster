package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite", "./data.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS temperatures (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		value REAL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(query)
	return err
}

func SaveReading(temp float64) error {
	_, err := DB.Exec("INSERT INTO temperatures (value) VALUES (?)", temp)
	return err
}
