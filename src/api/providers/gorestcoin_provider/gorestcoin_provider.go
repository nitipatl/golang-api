package gorest_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	restclient "github.com/nitipatl/golang-api/src/api/clients/rest_client"
	"github.com/nitipatl/golang-api/src/api/domain/gorestcoin"
)

const (
	authHeaderName   = "Authorization"
	authHeaderFormat = "Bearer %s"
	urlCreateUser    = "https://gorest.co.in/public-api/users"
)

func getAuthHeader(accessToken string) string {
	return fmt.Sprintf(authHeaderFormat, accessToken)
}
func CreateUser(accessToken string, request gorestcoin.CreateUserRequest) (*gorestcoin.CreateUserResponse, gorestcoin.GorestError) {
	headerToken := getAuthHeader(accessToken)
	headers := http.Header{}
	headers.Set(authHeaderName, headerToken)
	response, err := restclient.Post(urlCreateUser, request, headers)
	if err != nil {
		log.Println("error api:", err)
		return nil, gorestcoin.GorestError{
			Code: http.StatusInternalServerError,
			Data: nil,
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("error body api:", err)
		return nil, gorestcoin.GorestError{
			Code: int64(response.StatusCode),
			Data: nil,
		}
	}

	defer response.Body.Close()

	var result gorestcoin.StatusCodeResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println("error unmarshal api:", err)
		return nil, gorestcoin.GorestError{
			Code: int64(result.Code),
			Data: nil,
		}
	}

	if result.Code == 422 {
		var resultErr gorestcoin.GorestError
		if err := json.Unmarshal(bytes, &resultErr); err != nil {
			return nil, gorestcoin.GorestError{
				Code: int64(result.Code),
				Data: nil,
			}
		}
		return nil, resultErr
	} else if result.Code > 299 {
		return nil, gorestcoin.GorestError{
			Code: int64(response.StatusCode),
			Data: nil,
		}
	}
	var resultUser gorestcoin.CreateUserResponse
	if err := json.Unmarshal(bytes, &resultUser); err != nil {
		log.Println("error unmarshal api:", err)
		return nil, gorestcoin.GorestError{
			Code: int64(result.Code),
			Data: nil,
		}
	}
	return &resultUser, gorestcoin.GorestError{}
}
