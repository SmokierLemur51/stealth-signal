package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Handler struct {
	Port string
	DB   *sqlx.DB
}

func CreateHandler(router *http.ServeMux) *Handler {
	h := &Handler{Port: ":5000"}
	h.ConnectDatabasePSQL()
	h.RegisterRoutes(router)
	return h
}

func (h *Handler) ConnectDatabaseSQLITE(file string) {
	var err error
	h.DB, err = sqlx.Connect("sqlite3", fmt.Sprintf("instance/%s.db", file))
	if err != nil {
		panic(err)
	}
}

func (h *Handler) ConnectDatabasePSQL() {
	var e error
	h.DB, e = sqlx.Connect(
		"postgres", fmt.Sprintf("user=%s password=%s dbname=%s",
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASS"),
			os.Getenv("PG_DATABASE")))
	if e != nil {
		panic(e)
	}
}
