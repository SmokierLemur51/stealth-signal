package data

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSchema(db *sqlx.DB) {
	schema := `
		CREATE TABLE IF NOT EXISTS groups (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			group_name VARCHAR(150) UNIQUE, 
			secret_phrase_hash TEXT
		);

		CREATE TABLE IF NOT EXISTS group_members (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			group_id INTEGER, 
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			codename VARCHAR(60),
			public_key TEXT UNIQUE
		);

		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			secret_phrase_hash TEXT, 
			message TEXT
		)
	`
	db.MustExec(schema)
}
