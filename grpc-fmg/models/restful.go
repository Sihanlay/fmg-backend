package models

type RestfulAPIResult struct {
	Status bool `json:"status"`
	ErrCode int `json:"errcode"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

