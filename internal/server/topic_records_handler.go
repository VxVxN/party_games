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
	Topics   []string `json:"topics"`
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
}

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

	if req.PageSize == 0 {
		req.PageSize = 10
	}

	if len(req.Topics) == 0 {
		ErrResponse(w, http.StatusBadRequest, fmt.Errorf("topics is empty"))
		return
	}
	lines, err := getRecordsByTopics(req.Topics)
	if err != nil {
		ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	result := GetDataPage(lines, req.Page, req.PageSize)
	if result == nil {
		result = make([]string, 0)
	}
	SuccessResponse(w, TopicRecordsResponse{
		Records:   result,
		CountPage: len(lines)/req.PageSize + 1,
	})
}

func getRecordsByTopics(topics []string) ([]string, error) {
	if isAllTopic(topics) {
		files, err := os.ReadDir("topics")
		if err != nil {
			return nil, err
		}
		topics = []string{}
		for _, file := range files {
			topics = append(topics, file.Name())
		}
	}
	lines, err := readRecordsByTopics(topics)
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func readRecordsByTopics(topics []string) ([]string, error) {
	var lines []string
	for _, topic := range topics {
		data, err := os.ReadFile("topics/" + topic)
		if err != nil {
			return nil, err
		}
		lines = append(lines, strings.Split(string(data), "\n")...)
	}
	return lines, nil
}

func isAllTopic(topics []string) bool {
	for _, topic := range topics {
		if topic == "all" {
			return true
		}
	}
	return false
}

func UnmarshalRequest(body io.ReadCloser, reqStruct interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(&reqStruct)
}

func GetDataPage(data []string, page, pageSize int) []string {
	start := (page - 1) * pageSize
	stop := start + pageSize

	if start > len(data) {
		return nil
	}

	if stop > len(data) {
		stop = len(data)
	}
	return data[start:stop]
}
