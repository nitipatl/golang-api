package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	GetTest(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, `{"status":"OK!"}`, response.Body.String())
}
