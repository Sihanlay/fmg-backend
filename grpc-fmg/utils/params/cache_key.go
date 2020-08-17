package paramsUtils

import (
	"strconv"
)


// 构建缓存key
func CacheBuildKey(model string, keys ...interface{}) string {
	pay := model

	for _, key := range keys {
		pay += ":"
		switch k := key.(type) {
		case int:
			pay += strconv.Itoa(k)
		case int64:
			pay += strconv.Itoa(int(k))
		case int16:
			pay += strconv.Itoa(int(k))
		case string:
			pay += k
		}
	}
	return pay
}
