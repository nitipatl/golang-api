package users

type CreateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Gender string `json:"gender"`
}

type CreateUserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Gender string `json:"gender"`
}
