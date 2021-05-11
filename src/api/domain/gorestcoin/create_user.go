package gorestcoin

type CreateUserRequest UserData

type CreateUserResponse struct {
	Code int64    `json:"code"`
	Data UserData `json:"data"`
}

type UserData struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
