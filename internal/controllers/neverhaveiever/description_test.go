package neverhaveiever

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_DescriptionHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	controller := &Controller{}
	controller.DescriptionHandler(w, r)

	res := w.Result()
	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	actual := string(data)
	actual = strings.ReplaceAll(actual, "\\n", "\n")
	assert.Equal(t, fmt.Sprintf("{\"result\":\"%s\"}", descriptionText), actual)
}
