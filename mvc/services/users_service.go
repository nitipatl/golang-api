package services

import (
	"github.com/nitipatl/golang-api/mvc/domain"
	"github.com/nitipatl/golang-api/mvc/utils"
)

type usersService struct {
}

var (
	UserService *usersService
)

func (u *usersService) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
