package main

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger/v2"

	_ "github.com/VxVxN/party_games/docs"
	"github.com/VxVxN/party_games/internal/server"
)

//	@title		Party games API
//	@version	1.0

//	@host	localhost:8080

func main() {
	server := server.NewServer()

	router := chi.NewRouter()

	router.Get("/swagger/*", httpSwagger.Handler())

	router.Get("/neverhaveiever/description", server.LogMiddleware(server.INeverController.DescriptionHandler))
	router.Get("/neverhaveiever/topic/list", server.LogMiddleware(server.INeverController.TopicListHandler))
	router.Post("/neverhaveiever/topic/records", server.LogMiddleware(server.INeverController.TopicRecordsHandler))

	if err := server.ListenAndServe(router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
