package main

import (
	"log"
	"net/http"

	"github.com/VxVxN/party_games/internal/server"
)

func main() {
	server := server.NewServer()

	router := http.NewServeMux()
	router.HandleFunc("GET /topic/list", server.TopicListHandler)
	router.HandleFunc("POST /topic/records", server.TopicRecordsHandler)

	if err := server.ListenAndServe(router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
