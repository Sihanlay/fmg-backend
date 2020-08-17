package elasticsearch

//import (
//	"bytes"
//	productException "wzlz-backend/exceptions/product"
//	"wzlz-backend/utils"
//	"wzlz-backend/utils/log"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"strings"
//)
//
//// 索引
//func Create(_index, _type string, _id interface{}, body *bytes.Buffer) {
//	response, err := utils.Requests(
//		"PUT",
//		fmt.Sprintf(
//			"http://%s:%d/%s/%s/%v",
//			utils.GlobalConfig.Elasticsearch.Host,
//			utils.GlobalConfig.Elasticsearch.Port,
//			_index, _type, _id),
//		body)
//	log(response, err)
//}
//
//// 更新
//func Update(_index, _type string, _id interface{}, body *bytes.Buffer) {
//	response, err := utils.Requests(
//		"POST",
//		fmt.Sprintf(
//			"http://%s:%d/%s/%s/%v/_update",
//			utils.GlobalConfig.Elasticsearch.Host,
//			utils.GlobalConfig.Elasticsearch.Port,
//			_index, _type, _id),
//		body)
//	log(response, err)
//}
//
//// 删除
//func Delete(_index, _type string, _id interface{}) {
//	response, err := utils.Requests(
//		"DELETE",
//		fmt.Sprintf(
//			"http://%s:%d/%s/%s/%v",
//			utils.GlobalConfig.Elasticsearch.Host,
//			utils.GlobalConfig.Elasticsearch.Port,
//			_index, _type, _id),
//		nil)
//	log(response, err)
//}
//
//// 全局搜索
//func GlobalSearch(_index, key string, conf ...Config) map[string]interface{} {
//	payload := newGlobalSearchBody(key, conf...)
//	response, _ := utils.Requests(
//		"GET",
//		fmt.Sprintf(
//			"http://%s:%d/%s/_search?pretty",
//			utils.GlobalConfig.Elasticsearch.Host,
//			utils.GlobalConfig.Elasticsearch.Port,
//			_index),
//		payload)
//
//	if response.StatusCode != http.StatusOK {
//		panic(productException.SearchFail())
//	}
//
//	body := response.Body
//	defer body.Close()
//	// 格式化数据
//	if b, err := ioutil.ReadAll(body); err == nil {
//		var data map[string]interface{}
//		result := make(map[string]interface{})
//
//		if err := json.Unmarshal(b, &data); err == nil {
//			hits := data["hits"].(map[string]interface{})
//			// 获取检索条数
//			result["total"] = hits["total"].(map[string]interface{})["value"].(float64)
//			result["result"] = make([]map[string]interface{}, 0, int(result["total"].(float64)))
//			// 格式化每条数据
//			for index, hit := range hits["hits"].([]interface{}) {
//				// 本体数据
//				result["result"] = append(result["result"].([]map[string]interface{}), hit.(map[string]interface{})["_source"].(map[string]interface{}))
//				// 替换高亮数据
//				if v, ok := hit.(map[string]interface{})["highlight"].(map[string]interface{}); ok {
//					for key, value := range v {
//						// 替换字符串类型
//						target := result["result"].([]map[string]interface{})[index][key]
//						if t, ok := target.(string); ok {
//							// 每条高亮语句逐一替换
//							for _, k := range value.([]interface{}) {
//								lk := strings.Replace(k.(string), "</em>", "", -1)
//								lk = strings.Replace(lk, "<em>", "", -1)
//								t = strings.Replace(t, lk, k.(string), -1)
//							}
//							// 替换原数据
//							result["result"].([]map[string]interface{})[index][key] = t
//						}
//					}
//				}
//			}
//			if len(conf) > 0 {
//				if conf[0].size == 0 {
//					result["limit"] = 10
//				} else {
//					result["limit"] = conf[0].size
//				}
//				result["page"] = (conf[0].from / int(result["limit"].(int))) + 1
//			} else {
//				result["page"] = 1
//				result["limit"] = 10
//			}
//			return result
//		}
//	}
//	panic(productException.SearchFail())
//}
//
//// 日志
//func log(response *http.Response, err error) {
//	defer func() {
//		if err := recover(); err != nil {
//			logUtils.Println("索引失败", err)
//		}
//	}()
//	if err != nil {
//		b, _ := ioutil.ReadAll(response.Body)
//		logUtils.Println("索引失败", string(b), err)
//	}
//}
