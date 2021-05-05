package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nitipatl/golang-api/mvc/services"
	"github.com/nitipatl/golang-api/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, appErr := services.UserService.GetUser(uint64(userId))
	if appErr != nil {
		jsonValue, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
