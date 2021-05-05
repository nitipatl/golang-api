package app

import (
	"github.com/nitipatl/golang-api/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
