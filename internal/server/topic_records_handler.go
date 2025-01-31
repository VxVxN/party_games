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
	Topics []string `json:"topics"`
	Page   int      `json:"page"`
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

	if len(req.Topics) == 0 {
		ErrResponse(w, http.StatusBadRequest, fmt.Errorf("topics is empty"))
		return
	}

	var isAllTopic bool
	for _, topic := range req.Topics {
		if topic == "all" {
			isAllTopic = true
			break
		}
	}

	topics := req.Topics
	if isAllTopic {
		files, err := os.ReadDir("topics")
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		for _, file := range files {
			topics = append(topics, file.Name())
		}
	}
	var lines []string
	for _, topic := range topics {
		data, err := os.ReadFile("topics/" + topic)
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		lines = append(lines, strings.Split(string(data), "\n")...)
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
