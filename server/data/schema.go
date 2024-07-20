package data

import (
	"github.com/jmoiron/sqlx"
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

func CreateSchemaPSQL(db *sqlx.DB) {
	schema := `
		CREATE TABLE IF NOT EXISTS groups (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			group_name VARCHAR(150) UNIQUE, 
			secret_phrase_hash VARCHAR(1500)
		);

		CREATE TABLE IF NOT EXISTS group_members (
			id SERIAL PRIMARY KEY,
			group_id INTEGER, 
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			codename VARCHAR(60),
			public_key VARCHAR(1500) UNIQUE
		);

		CREATE TABLE IF NOT EXISTS messages (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			secret_phrase_hash VARCHAR(1500), 
			message VARCHAR(1500)
		)
	`
	db.MustExec(schema)
}
