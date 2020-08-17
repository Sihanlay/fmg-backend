package elasticsearch

import (
	"bytes"
	"encoding/json"
)

type dict map[string]interface{}

// 新建全局搜索请求
func newGlobalSearchBody(key string, conf ...Config) *bytes.Buffer {
	body := dict {
		"query": dict {
			"bool": dict {},
		},
		"highlight": newHighLight(),
	}
	// 添加语句检索
	if len(key) > 0 {
		body["query"].(dict)["bool"].(dict)["must"] = []dict {
			{
				"multi_match": dict {
					"query": key,
				},
			},
		}
	}

	if len(conf) > 0 {
		// 分页配置
		body["from"] = conf[0].from
		if conf[0].size == 0 {
			body["size"] = 10
		} else {
			body["size"] = conf[0].size
		}
		// 过滤条件
		filter := make([]dict, 0, len(conf[0].filter))
		for i, v := range conf[0].filter {
			filter = append(filter, dict {
				"terms": dict {
					i: v,
				},
			})
		}
		body["query"].(dict)["bool"].(dict)["filter"] = filter
		// 字段展示过滤
		body["_source"] = conf[0].source
		// 字段匹配
		if len(key) > 0 {
			body["query"].(dict)["bool"].(dict)["must"].([]dict)[0]["multi_match"].(dict)["fields"] = conf[0].fields
		}
	}

	if re, err := json.Marshal(body); err == nil {
		return bytes.NewBuffer(re)
	}
	return nil
}

// 新建高亮请求体信息
func newHighLight() dict {
	return dict {
		"require_field_match": false,
		"fields": dict {
			"*": dict{},
		},
	}
}
