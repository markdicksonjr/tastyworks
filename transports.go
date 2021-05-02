package tastyworks

import "net/http"

type ErrorResult struct {
	Code    string
	Message string
}

func SetStandardHeadersOnRequest(req *http.Request) {
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,tr;q=0.8")
	req.Header.Set("Accept-Version", "v1")
	req.Header.Set("Authorization", "null")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "api.tastyworks.com")
	req.Header.Set("Origin", "https://trade.tastyworks.com")
	req.Header.Set("Referer", "https://trade.tastyworks.com/tw")
}