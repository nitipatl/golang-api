package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMock = false
	mocks      = make(map[string]*Mock)
)

func StartMock() {
	enableMock = true
}

func StopMock() {
	enableMock = false
}

type Mock struct {
	Url    string
	Method string
	Resp   *http.Response
	Err    error
}

func AddMock(mock Mock) {
	mocks[mock.Url] = &mock
}

func FlushMock() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	if enableMock {
		mock := mocks[url]
		if mock == nil {
			return nil, errors.New("no mockup found")
		}
		return mock.Resp, mock.Err
	}

	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonBytes))
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))

	if err != nil {
		return nil, err
	}

	request.Header = headers
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	return client.Do(request)
}
