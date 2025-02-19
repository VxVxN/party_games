package neverhaveiever

import (
	"net/http"
	"os"

	"github.com/VxVxN/party_games/pkg/httptools"
)

func (controller *Controller) TopicListHandler(w http.ResponseWriter, r *http.Request) {
	result, err := getTopicList()
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	httptools.SuccessResponse(w, result)
}

func getTopicList() ([]string, error) {
	result := []string{"all"}
	dirEntities, err := os.ReadDir(topicsPath)
	if err != nil {
		return nil, err
	}
	for _, entity := range dirEntities {
		result = append(result, entity.Name())
	}
	return result, nil
}
