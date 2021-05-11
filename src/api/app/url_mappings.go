package app

import (
	"github.com/nitipatl/golang-api/src/api/controllers/test"
	"github.com/nitipatl/golang-api/src/api/controllers/users"
)

func mapUrls() {
	router.GET("/", test.GetTest)
	router.POST("/users", users.CreateUser)
}
