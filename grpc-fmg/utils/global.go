package utils

import (
	"encoding/json"
	"grpc-demo/config"
	"grpc-demo/models"

)

var GlobalConfig models.SystemConfiguration
var tempConfig models.SystemConfiguration

func InitGlobal() {
	json.Unmarshal(config.Config(),&GlobalConfig)
}
