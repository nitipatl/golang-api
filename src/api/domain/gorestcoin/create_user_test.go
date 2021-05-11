package gorestcoin

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserRequestAsJson(t *testing.T) {
	request := CreateUserRequest{
		Name:   "Nitipat L",
		Email:  "iamsvz+junk2@gmail.com",
		Gender: "Male",
		Status: "Active",
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t,
		`{"name":"Nitipat L","email":"iamsvz+junk2@gmail.com","gender":"Male","status":"Active"}`,
		string(bytes))

	var target CreateUserRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, "Nitipat L")
	assert.EqualValues(t, target.Email, "iamsvz+junk2@gmail.com")
	assert.EqualValues(t, target.Gender, "Male")
	assert.EqualValues(t, target.Status, "Active")
}

func TestCreateUserResponseAsJson(t *testing.T) {
	response := CreateUserResponse{
		Code: 200,
		Data: UserData{
			Name:   "Nitipat L",
			Email:  "iamsvz+junk2@gmail.com",
			Gender: "Male",
			Status: "Active",
		},
	}
	bytes, err := json.Marshal(response)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t,
		`{"code":200,"data":{"name":"Nitipat L","email":"iamsvz+junk2@gmail.com","gender":"Male","status":"Active"}}`,
		string(bytes))

	var target CreateUserResponse
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Code, 200)
}

func TestCreateUserResponseErrorAsJson(t *testing.T) {
	response := GorestError{
		Code: 422,
		Data: []GorestFieldsError{
			{
				Field:   "email",
				Message: "has already been taken",
			},
		},
	}
	bytes, err := json.Marshal(response)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t,
		`{"code":422,"data":[{"field":"email","message":"has already been taken"}]}`,
		string(bytes))

	var target GorestError
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Code, 422)
}
