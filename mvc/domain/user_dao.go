package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nitipatl/golang-api/mvc/utils"
)

type userDao struct{}

type usersServiceInterface interface {
	GetUser(uint64) (*User, *utils.ApplicationError)
}

var (
	users = map[uint64]*User{
		123: &User{Id: 123, FirstName: "Nitipat", LastName: "L", Email: "iamsvz@gmail.com"},
	}
	UserDao usersServiceInterface
)

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userId uint64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("Not found user %v!", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
