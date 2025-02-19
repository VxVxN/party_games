package neverhaveiever

import (
	"github.com/rs/zerolog"
)

const topicsPath = "games/never_have_i_ever/topics"

type Controller struct {
	fileDataByName map[string][]byte
	logger         zerolog.Logger
}

func NewController(logger zerolog.Logger) *Controller {
	return &Controller{
		fileDataByName: make(map[string][]byte),
		logger:         logger,
	}
}
