package server

import (
	"encoding/json"
	"net/http"
	"os"
)

type errorResponse struct {
	Error string `json:"error"`
}

func (server *Server) TopicListHandler(w http.ResponseWriter, r *http.Request) {
	result, err := getTopicList(w)
	if err != nil {
		ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(w, result)
}

func getTopicList(w http.ResponseWriter) ([]string, error) {
	result := []string{"all"}
	dirEntities, err := os.ReadDir("topics")
	if err != nil {
		return nil, err
	}
	for _, entity := range dirEntities {
		result = append(result, entity.Name())
	}
	return result, nil
}

func ErrResponse(w http.ResponseWriter, status int, err error) {
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE, PUT")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	data, err := json.Marshal(errorResponse{Error: err.Error()})
	if err != nil {
		// todo
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func SuccessResponse(w http.ResponseWriter, resp interface{}) {
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE, PUT")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if resp == nil {
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		// todo
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
