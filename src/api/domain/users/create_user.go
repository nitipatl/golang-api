package users

type CreateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Active string `json:"active"`
	Gender string `json:"gender"`
}

type CreateUserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Active string `json:"active"`
	Gender string `json:"gender"`
}
