package services

import (
	"net/http"

	"github.com/nitipatl/golang-api/mvc/domain"
	"github.com/nitipatl/golang-api/mvc/utils"
)

type itemsService struct{}

var (
	ItemService itemsService
)

func (t *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement message",
		StatusCode: http.StatusInternalServerError,
	}
}
