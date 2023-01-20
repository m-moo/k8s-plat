package helm

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetChartsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	GetChartsHandler(c)

	assert.Equal(t, 200, w.Code)

	var got []chartList
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}
}
