package requestsUtils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const UrlBase = "http://internal_host/api_internal/"

func Do(mothod, url string, body []byte) (map[string]interface{}, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	request, err := http.NewRequest(mothod, UrlBase+url, reader)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	v, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(v, &data); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
