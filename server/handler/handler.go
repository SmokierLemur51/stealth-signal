package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SmokierLemur51/stealth-signal-server/data"
	"github.com/SmokierLemur51/stealth-signal-server/util"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Handler struct {
	Port string
	DB   *sqlx.DB
}

func CreateHandler(router *http.ServeMux) Handler {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	h := Handler{Port: ":5000"}
	h.ConnectDatabase(os.Getenv("SQLITE_DB_FILE"))
	h.RegisterRoutes(router)
	return h
}

func (h *Handler) ConnectDatabase(file string) {
	var err error
	h.DB, err = sqlx.Open("sqlite3", fmt.Sprintf("instance/%s.db", file))
	if err != nil {
		panic(err)
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", h.IndexHandler)
	router.HandleFunc("GET /test", h.TestHandler)
	router.HandleFunc("GET /group/{group_name}/latest", h.LatestHandler)
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (h *Handler) TestHandler(w http.ResponseWriter, r *http.Request) {
	message, err := util.HashMessage("This is a test")
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusBadRequest)
	}
	w.Write([]byte(message))
}

func (h *Handler) LatestHandler(w http.ResponseWriter, r *http.Request) {
	// Check for group
	group := data.Group{}
	h.DB.Get(&group, "SELECT * FROM groups WHERE group_name=?", r.PathValue("group_name"))
}
