package wechat


type WeCharClient struct {
	Appid            string   `json:"appid"`               // 应用唯一标识，在微信开放平台提交应用审核通过后获得
	Secret           string   `json:"secret"`              // 应用密钥AppSecret，在微信开放平台提交应用审核通过后获得
	RedirectUri      string   `json:"redirect_uri"`        // 回调地址
	Scope            []string `json:"scope"`               // 应用授权作用域，拥有多个作用域用逗号（,）分隔，网页应用目前仅填写snsapi_login即
	State            string   `json:"state"`               // 用于保持请求和回调的状态，授权请求后原样带回给第三方。该参数可用于防止csrf攻击（跨站请求伪造攻击），建议第三方带上该参数，可设置为简单的随机数加session进行校验
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`  // 接口调用凭证
	ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	OpenId       string `json:"open_id"`       // 授权用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	Unionid      string `json:"unionid"`       // 当且仅当该网站应用已获得该用户的userinfo授权时，才会出现该字段。
}

type UserInfo struct {
	Openid     string   `json:"openid"`        // 普通用户的标识，对当前开发者帐号唯一
	Nickname   string   `json:"nickname"`      // 普通用户昵称
	Sex        int      `json:"sex"`           // 普通用户性别，1为男性，2为女性
	Province   string   `json:"province"`      // 普通用户个人资料填写的省份
	City       string   `json:"city"`          // 普通用户个人资料填写的城市
	Country    string   `json:"country"`       // 国家，如中国为CN
	Headimgurl string   `json:"headimgurl"`    // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
	Privilege  []string `json:"privilege"    ` // 用户特权信息，json数组，如微信沃卡用户为（chinaunicom）
	Unionid    string   `json:"unionid"`       // 用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的unionid是唯一的。
}

