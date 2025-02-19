package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/VxVxN/party_games/internal/controllers/neverhaveiever"
)

type Server struct {
	INeverController *neverhaveiever.Controller
	logger           zerolog.Logger
}

func NewServer() *Server {
	logger := log.Logger
	return &Server{
		logger:           logger,
		INeverController: neverhaveiever.NewController(logger),
	}
}

func (server *Server) ListenAndServe(handler http.Handler) error {
	port := 8080
	server.logger.Info().Int("port", port).Msg("Starting server")
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}

func (server *Server) LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		server.logger.Info().Str("method", r.Method).Str("path", r.URL.Path).Msg("Received request")
		next.ServeHTTP(w, r)
		server.logger.Info().Str("method", r.Method).Str("path", r.URL.Path).Dur("duration", time.Now().Sub(startTime)).Msg("Completed request")
	}
}
