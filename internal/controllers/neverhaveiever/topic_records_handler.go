package neverhaveiever

import (
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/VxVxN/party_games/pkg/httptools"
)

type TopicRecordsRequest struct {
	Topics   []string `json:"topics"`
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
	Refresh  bool     `json:"refresh"`
}

type TopicRecordsResponse struct {
	Records   []string `json:"records"`
	CountPage int      `json:"count_page"`
}

func (controller *Controller) TopicRecordsHandler(w http.ResponseWriter, r *http.Request) {
	var req TopicRecordsRequest

	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal request body: %v", err))
		return
	}

	if req.PageSize == 0 {
		req.PageSize = 10
	}

	if len(req.Topics) == 0 {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("topics is empty"))
		return
	}
	lines, err := controller.getRecordsByTopics(req.Topics)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	result := GetDataPage(lines, req.Page, req.PageSize)
	if result == nil {
		result = make([]string, 0)
	}
	if req.Refresh {
		rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	}
	httptools.SuccessResponse(w, TopicRecordsResponse{
		Records:   result,
		CountPage: int(math.Ceil(float64(len(lines)) / float64(req.PageSize))),
	})
}

func (controller *Controller) getRecordsByTopics(topics []string) ([]string, error) {
	if isAllTopic(topics) {
		files, err := os.ReadDir(topicsPath)
		if err != nil {
			return nil, err
		}
		topics = []string{}
		for _, file := range files {
			topics = append(topics, file.Name())
		}
	}
	lines, err := controller.readRecordsByTopics(topics)
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func (controller *Controller) readRecordsByTopics(topics []string) ([]string, error) {
	var lines []string
	for _, topic := range topics {
		var data []byte
		var ok bool
		data, ok = controller.fileDataByName[topic]
		if !ok {
			var err error
			data, err = os.ReadFile(path.Join(topicsPath, topic))
			if err != nil {
				return nil, err
			}
			controller.fileDataByName[topic] = data
		}
		if data == nil {
			continue
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
