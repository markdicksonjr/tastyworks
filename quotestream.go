package tastyworks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetQuoteStreamerTokenResult struct {
	Data GetQuoteStreamerTokenResultData
	Context string
}

type GetQuoteStreamerTokenResultData struct {
	Token        string `json:"token"`
	StreamerUrl  string `json:"streamer-url"`
	WebsocketUrl string `json:"websocket-url"`
	Level        string `json:"level"`
}

func GetQuoteStreamerToken(sessionToken string) (GetQuoteStreamerTokenResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/quote-streamer-tokens", nil)

	if err != nil {
		return GetQuoteStreamerTokenResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetQuoteStreamerTokenResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetQuoteStreamerTokenResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetQuoteStreamerTokenResult{}, err
	}

	v := GetQuoteStreamerTokenResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetQuoteStreamerTokenResult{}, err
	}

	return v, nil
}
