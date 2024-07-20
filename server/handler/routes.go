package handler

import (
	"net/http"

	"github.com/SmokierLemur51/stealth-signal-server/data"
	"github.com/SmokierLemur51/stealth-signal-server/util"
)

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
