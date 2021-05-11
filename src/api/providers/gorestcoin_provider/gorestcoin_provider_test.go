package gorest_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	restclient "github.com/nitipatl/golang-api/src/api/clients/rest_client"
	"github.com/nitipatl/golang-api/src/api/domain/gorestcoin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMock()
	os.Exit(m.Run())
}

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("{{test-token}}")
	assert.EqualValues(t, "Bearer {{test-token}}", header)
}

func TestCreateUserFailed(t *testing.T) {
	restclient.FlushMock()
	restclient.AddMock(restclient.Mock{
		Url:    "https://gorest.co.in/public-api/users",
		Method: http.MethodPost,
		Err:    errors.New("error from api"),
	})
	user, err := CreateUser("de280ea2a0b7a56cd3694f919407258c786968350f8a101528134e7ffddc8435", gorestcoin.CreateUserRequest{
		Name:   "Nitipat L",
		Email:  "iamsvz+junk4@gmail.com",
		Gender: "Male",
		Status: "Active",
	})
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestCreateUserExistEmailFailed(t *testing.T) {
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
	user, err := CreateUser("de280ea2a0b7a56cd3694f919407258c786968350f8a101528134e7ffddc8435", gorestcoin.CreateUserRequest{
		Name:   "Nitipat L",
		Email:  "iamsvz+junk4@gmail.com",
		Gender: "Male",
		Status: "Active",
	})
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusUnprocessableEntity, err.Code)
}

func TestCreateUserSuccess(t *testing.T) {
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
	user, _ := CreateUser("", gorestcoin.CreateUserRequest{
		Name:   "Nitipat L",
		Email:  "iamsvz+junk4@gmail.com",
		Gender: "Male",
		Status: "Active",
	})
	assert.EqualValues(t, 201, user.Code)
	assert.EqualValues(t, "Nitipat L", user.Data.Name)
}
