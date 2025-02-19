package main

import (
	"log"
	"net/http"

	"github.com/VxVxN/party_games/internal/server"
)

func main() {
	server := server.NewServer()

	router := http.NewServeMux()
	router.Handle("GET /topic/list", server.LogMiddleware(server.INeverController.TopicListHandler))
	router.Handle("POST /topic/records", server.LogMiddleware(server.INeverController.TopicRecordsHandler))

	if err := server.ListenAndServe(router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
