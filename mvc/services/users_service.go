package services

import (
	"github.com/nitipatl/golang-api/mvc/domain"
	"github.com/nitipatl/golang-api/mvc/utils"
)

func GetUsert(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
