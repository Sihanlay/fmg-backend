package elasticsearch

// 搜索配置
type Config struct {
	// 分页配置
	from int
	size int
	// 显示字段
	source []string
	// 搜索字段
	fields []string
	// 过滤字段
	filter map[string]interface{}
}

// 新建产品搜索配置
func NewProductSearchConfig(from, size int) Config {
	return Config{
		from: from,
		size: size,
		fields: []string{"name", "description", "details", "additional"},
		source: []string{"id", "name", "update_time", "create_time",
			   			 "description", "cover", "status", "star"},
		filter: make(map[string]interface{}),
	}
}

// 添加过滤条件
func (c *Config) AddFilter(key string, value interface{}) {
	switch value.(type) {
	case []string:
		c.filter[key] = value
	default:
		c.filter[key] = []interface{}{value}
	}
}
