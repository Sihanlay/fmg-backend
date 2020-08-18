package constants

const (
	// ------ 系统 ------
	// cookie
	SystemCookie = "farm_SYSTEM_COOKIE"
	// session过期时间 10小时
	SessionExpires = 864000
	// cookie过期时间  7天
	CookieExpires = 3600 * 24 * 7
	// 登陆态session名称
	SessionName = "farm_AUTHENTICATION_KEY"
	// 登陆模式
	ApiMode = "api-mode"
	// token
	ApiToken = "api-token"

	// ------ 资源 ------
	StorageTokenTime = 300

	// 任务队列长度
	QueueTaskLength = 100
)
