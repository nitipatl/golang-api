package domain

import (
	"fmt"
	"net/http"

	"github.com/nitipatl/golang-api/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Nitipat", LastName: "L", Email: "iamsvz@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("Not found user %v!", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
