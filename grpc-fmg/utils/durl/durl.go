package durl

import (
	paramsException "grpc-demo/exceptions/params"
	"encoding/base64"
	"regexp"
)

func DataUrlParser(durl string) []byte {
	re := regexp.MustCompile(`^data:(.*?);base64,(.*?)$`)
	r := re.FindSubmatch([]byte(durl))

	if len(r) != 3 {
		panic(paramsException.DataUrlParserFail())
	}
	data, err := base64.StdEncoding.DecodeString(string(r[2][:]))
	if err != nil {
		panic(paramsException.DataUrlParserFail())
	}
	return data
}
