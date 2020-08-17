package constants

const (
	DataRoot = "../data"

	// 账户
	Account = DataRoot + "/account"

	// 产品
	//Product = DataRoot + "/product"

	// ----- 子模块 -----

	// 账户头像
	AccountAvator = Account + "/avator"

	// 产品封面
	//ProductCover = Product + "/cover"

	// nginx静态资源映射
	NginxResourcePath = "/resource_internal"
)

var StorageMapping = map[string]string{
	"account_avator": AccountAvator,
	//"product_cover":  ProductCover,
}

var MimeToExtMapping = map[string]string{
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"bmp":  "image/bmp",
	"png":  "image/png",
	"gif":  "image/gif",
	"svg":  "image/svg",
}
