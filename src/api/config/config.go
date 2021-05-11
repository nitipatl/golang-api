package config

import "os"

const (
	apiGorestToken = "GOREST_SECRET_TOKEN"
)

var (
	gorestToken = os.Getenv(apiGorestToken)
)

func GetGorestToken() string {
	return gorestToken
}
