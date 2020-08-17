package utils

import (
	"io"
	"net/http"
)

func Requests(method string, url string, body io.Reader) (*http.Response, error){
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, body)
	request.Header.Set("Content-Type", "application/json")
	return client.Do(request)
}
