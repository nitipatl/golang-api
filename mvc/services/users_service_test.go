package services

import (
	"log"
	"net/http"
	"testing"

	"github.com/nitipatl/golang-api/mvc/domain"
	"github.com/nitipatl/golang-api/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock
	getUserFunc func(userId uint64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func (u *usersDaoMock) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return getUserFunc(userId)
}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func TestUserNotFoundDatabase(t *testing.T) {
	getUserFunc = func(userId uint64) (*domain.User, *utils.ApplicationError) {
		log.Println("Mock database...")
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "Not found user 0!",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not found user 0!", err.Message)

}

func TestUserFoundDatabase(t *testing.T) {
	getUserFunc = func(userId uint64) (*domain.User, *utils.ApplicationError) {
		log.Println("Mock database...")
		return &domain.User{
			Id:        userId,
			FirstName: "Nitipat2",
			LastName:  "L2",
			Email:     "iamsvz2@gmail.com",
		}, nil
	}
	user, err := UserService.GetUser(1234)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1234, user.Id)
	assert.EqualValues(t, "Nitipat2", user.FirstName)
	assert.EqualValues(t, "L2", user.LastName)
	assert.EqualValues(t, "iamsvz2@gmail.com", user.Email)

}
