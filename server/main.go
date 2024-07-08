package main

import (
	"log"
	"net/http"

	"github.com/SmokierLemur51/stealth-signal-server/data"
	"github.com/SmokierLemur51/stealth-signal-server/handler"
)

func main() {
	router := http.NewServeMux()
	handler := handler.CreateHandler(router)

	data.CreateSchema(handler.DB)
	data.InsertTestGroups(handler.DB)

	server := http.Server{
		Addr:    handler.Port,
		Handler: router,
	}

	log.Println("Startin server, listening on port " + handler.Port)
	log.Fatal(server.ListenAndServe())
}
