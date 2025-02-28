package neverhaveiever

import (
	"net/http"
	"os"

	"github.com/VxVxN/party_games/pkg/httptools"
)

// TopicListHandler Return list of topics
//
//	@Summary		Topic list
//	@Description	get topic list
//	@Tags			neverhaveiever
//	@Produce		json
//	@Failure		500	{object}	httptools.JsonErrorResponse
//	@Success		200	{object}	httptools.JsonSuccessResponse{result=[]string}	"desc"
//	@Router			/neverhaveiever/topic/list [get]
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
