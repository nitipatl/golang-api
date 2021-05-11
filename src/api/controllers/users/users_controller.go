package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitipatl/golang-api/src/api/domain/users"
	"github.com/nitipatl/golang-api/src/api/services"
	"github.com/nitipatl/golang-api/src/api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var request users.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.UserService.CreateUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
