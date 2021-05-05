package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nitipatl/golang-api/mvc/services"
	"github.com/nitipatl/golang-api/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, appErr)
		return
	}

	user, appErr := services.UserService.GetUser(uint64(userId))
	if appErr != nil {
		utils.RespondError(c, appErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
