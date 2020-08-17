package models


type FileFormat struct {

	Id int `json:"id"`

	Name string `json:"name"`

	Size int64 `json:"size"`

	Storage string `json:"storage"`

	CreateTime int64 `json:"create_time"`
}
