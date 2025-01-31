package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type TopicRecordsRequest struct {
	Topic string `json:"topic"`
	Page  int    `json:"page"`
}

const (
	itemsPerPage = 10
)

type TopicRecordsResponse struct {
	Records   []string `json:"records"`
	CountPage int      `json:"count_page"`
}

func (server *Server) TopicRecordsHandler(w http.ResponseWriter, r *http.Request) {
	var req TopicRecordsRequest

	if err := UnmarshalRequest(r.Body, &req); err != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal request body: %v", err))
		return
	}

	var lines []string
	if req.Topic == "all" {
		files, err := os.ReadDir("topics")
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		for _, file := range files {
			data, err := os.ReadFile("topics/" + file.Name())
			if err != nil {
				ErrResponse(w, http.StatusInternalServerError, err)
				return
			}
			lines = append(lines, strings.Split(string(data), "\n")...)
		}
	} else {
		data, err := os.ReadFile("topics/" + req.Topic)
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		lines = strings.Split(string(data), "\n")
	}
	SuccessResponse(w, TopicRecordsResponse{
		Records:   GetDataPage(lines, req.Page),
		CountPage: len(lines)/itemsPerPage + 1,
	})
}

func UnmarshalRequest(body io.ReadCloser, reqStruct interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(&reqStruct)
}

func GetDataPage(data []string, page int) []string {
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage

	if start > len(data) {
		return nil
	}

	if stop > len(data) {
		stop = len(data)
	}

	return data[start:stop]
}
