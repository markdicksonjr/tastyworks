package tastyworks

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type AuthorizationResult struct {
	Data  AuthorizationResultData
	Error ErrorResult
}

type AuthorizationResultData struct {
	User         AuthorizationResultUser
	SessionToken string `json:"session-token"`
}

type AuthorizationResultUser struct {
	Email      string
	Username   string
	ExternalId string `json:"external-id"`
}

func Authorize(username string, password string) (AuthorizationResult, error) {
	req, err := http.NewRequest("POST", TastyTradeHost+"/sessions", strings.NewReader(
		fmt.Sprintf("{\"login\":\"%s\",\"password\":\"%s\"}", username, password,
		)))

	if err != nil {
		return AuthorizationResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return AuthorizationResult{}, err
	}

	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return AuthorizationResult{}, err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)

	v := AuthorizationResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return AuthorizationResult{}, err
	}

	return v, nil
}
