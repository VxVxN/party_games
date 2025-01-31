package server

import (
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(":8080", handler)
}
