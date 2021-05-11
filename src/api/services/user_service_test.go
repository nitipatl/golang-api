package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	restclient "github.com/nitipatl/golang-api/src/api/clients/rest_client"
	"github.com/nitipatl/golang-api/src/api/domain/users"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMock()
	os.Exit(m.Run())
}

func TestCreateUserInvalidInputName(t *testing.T) {
	request := users.CreateUserRequest{}

	result, err := UserService.CreateUser(request)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Please input name, e-mail and gender", err.Message())
}

func TestCreateUserErrorFromGorest(t *testing.T) {
	restclient.FlushMock()
	json := `{"code": 422,"meta": null,"data": [{"field": "email","message": "has already been taken"}]}`
	restclient.AddMock(restclient.Mock{
		Url:    "https://gorest.co.in/public-api/users",
		Method: http.MethodPost,
		Resp: &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(json)),
		},
	})
	request := users.CreateUserRequest{
		Name:   "Nitipat L",
		Email:  "iamsvz+2@gmail.com",
		Gender: "Male",
	}

	result, err := UserService.CreateUser(request)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusUnprocessableEntity, err.Status())
	assert.EqualValues(t, "email:has already been taken ", err.Message())
}
