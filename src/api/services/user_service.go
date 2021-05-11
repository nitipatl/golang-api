package services

import (
	"strings"

	"github.com/nitipatl/golang-api/src/api/config"
	"github.com/nitipatl/golang-api/src/api/domain/gorestcoin"
	"github.com/nitipatl/golang-api/src/api/domain/users"
	gorest_provider "github.com/nitipatl/golang-api/src/api/providers/gorestcoin_provider"
	"github.com/nitipatl/golang-api/src/api/utils/errors"
)

type usersService struct{}

type userServiceInterface interface {
	CreateUser(users.CreateUserRequest) (*users.CreateUserResponse, errors.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &usersService{}
}

func (s *usersService) CreateUser(input users.CreateUserRequest) (*users.CreateUserResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)
	input.Gender = strings.TrimSpace(input.Gender)

	if input.Name == "" || input.Email == "" || input.Gender == "" {
		return nil, errors.NewBadRequestError("Please input name, e-mail and gender")
	}

	request := gorestcoin.CreateUserRequest{
		Name:   input.Name,
		Email:  input.Email,
		Status: "Active",
		Gender: input.Gender,
	}

	response, err := gorest_provider.CreateUser(config.GetGorestToken(), request)
	if err.Code == 422 {
		fieldsError := ""
		for _, errM := range err.Data {
			fieldsError += errM.Field + ":" + errM.Message + " "
		}
		return nil, errors.NewUpprocessRequestError(fieldsError)
	} else if err.Code > 0 {
		return nil, errors.NewInternalServerError("API 3rd party error something")
	}
	result := users.CreateUserResponse{
		Name:   response.Data.Name,
		Email:  response.Data.Email,
		Status: response.Data.Status,
		Gender: response.Data.Gender,
	}

	return &result, nil
}
