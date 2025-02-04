package server

import (
	"net/http"
)

type Server struct {
	fileDataByName map[string][]byte
}

func NewServer() *Server {
	return &Server{
		fileDataByName: make(map[string][]byte),
	}
}

func (server *Server) ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(":8080", handler)
}
