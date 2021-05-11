package users

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	restclient "github.com/nitipatl/golang-api/src/api/clients/rest_client"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMock()
	os.Exit(m.Run())
}

func TestCreateUserInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = &http.Request{}

	CreateUser(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, `{"status":400,"message":"invalid json body"}`, response.Body.String())
}

func TestSuccessCreateUser(t *testing.T) {
	restclient.FlushMock()
	json := `{"code": 201,"meta": null,"data": {"id": 1429,"name": "Nitipat L","email": "iamsvz+junk5@gmail.com","gender": "Male","status": "Active","created_at": "2021-05-11T09:03:01.504+05:30","updated_at": "2021-05-11T09:03:01.504+05:30"}}`
	restclient.AddMock(restclient.Mock{
		Url:    "https://gorest.co.in/public-api/users",
		Method: http.MethodPost,
		Resp: &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(json)),
		},
	})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "Nitipat L","email": "iamsvz+junk7@gmail.com","gender": "Male"}`))

	CreateUser(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, `{"name":"Nitipat L","email":"iamsvz+junk5@gmail.com","status":"Active","gender":"Male"}`, response.Body.String())
}
