package data

import (
	"database/sql"
	"time"

	"github.com/SmokierLemur51/stealth-signal-server/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Group struct {
	ID               int            `db:"id"`
	CreatedAt        time.Time      `db:"created_at"`
	GroupName        sql.NullString `db:"group_name"`
	SecretPhraseHash sql.NullString `db:"secret_phrase_hash"`
}

func InsertTestGroups(db *sqlx.DB) {
	stmt := "INSERT INTO groups (group_name, secret_phrase_hash) values (?, ?)"
	db.Exec(stmt, "frosty_fish", util.HashMessageNoError("mega supa secret"))
	db.Exec(stmt, "dumbledores_army", util.HashMessageNoError("we say his name"))
	db.Exec(stmt, "wolverines", util.HashMessageNoError("no more communists"))
}
