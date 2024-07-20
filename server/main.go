package main

import (
	"log"
	"net/http"

	"github.com/SmokierLemur51/stealth-signal-server/handler"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := http.NewServeMux()
	handler := handler.CreateHandler(router)

	// data.CreateSchemaPSQL(handler.DB)
	// data.InsertTestGroups(handler.DB)

	server := http.Server{
		Addr:    handler.Port,
		Handler: router,
	}

	log.Println("Startin server, listening on port " + handler.Port)
	log.Fatal(server.ListenAndServe())
}
