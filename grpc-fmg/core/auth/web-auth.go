package authbase

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"grpc-demo/constants"
	accountEnums "grpc-demo/enums/account"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	"grpc-demo/utils/hash"
	"time"
)

var sess = sessions.New(sessions.Config{
	Cookie: constants.SystemCookie,
})

type dAuthAuthorization struct {
	Account db.Account
	isLogin bool
	Context iris.Context
}

func (r *dAuthAuthorization) AccountModel() *db.Account {
	return &r.Account
}

func (r *dAuthAuthorization) CheckLogin() {
	if !r.isLogin {
		panic(accountException.AuthIsNotLogin())
	}
}

func (r *dAuthAuthorization) IsLogin() bool {
	return r.isLogin
}

func (r *dAuthAuthorization) IsAdmin() bool {
	return r.Account.Role == accountEnums.RoleAdmin
}

func (r *dAuthAuthorization) CheckAdmin() {
	r.CheckLogin()
	if !r.IsAdmin() {
		panic(accountException.NoPermission())
	}
}
func (r *dAuthAuthorization) LoadAuthenticationInfo() {
	r.loadFromCookie()
}

// 从session载入登录信息
func (r *dAuthAuthorization) loadFromSession() bool {
	// 阻止解码方法异常传递
	defer func() {
		recover()
	}()

	session := sess.Start(r.Context)
	sestring := session.GetString(constants.SessionName)
	if sestring == "" {
		return false
	}
	var sessionStruct cookieInfo
	hash.DecodeToken(sestring, &sessionStruct)

	if sessionStruct.ExpireTime <= time.Now().Unix() {
		return false
	}
	succ := r.fetchAccount(sessionStruct.AccountId)
	if succ {
		r.isLogin = true
	}
	return true
}

// 从cookie载入登录信息
func (r *dAuthAuthorization) loadFromCookie() bool {
	defer func() {
		recover()
	}()
	cookie := r.Context.GetCookie(constants.SessionName)
	if cookie == "" {
		return false
	}
	var cookieStruct cookieInfo
	hash.DecodeToken(cookie, &cookieStruct)
	if cookieStruct.ExpireTime <= time.Now().Unix() {
		return false
	}
	succ := r.fetchAccount(cookieStruct.AccountId)
	if succ {
		r.isLogin = true
	}
	return true
}

// 设置session
func (r *dAuthAuthorization) SetSession(aid int) {
	session := sess.Start(r.Context)

	if aid == 0 {
		session.Delete(constants.SessionName)
		return
	}
	payload := GenerateToken(aid, constants.SessionExpires)
	session.Set(constants.SessionName, payload)

}

// 设置cookie
func (r *dAuthAuthorization) SetCookie(aid int) string {
	payload := GenerateToken(aid, constants.CookieExpires)
	r.Context.SetCookieKV(constants.SessionName, payload)
	return ""
}

func (r *dAuthAuthorization) fetchAccount(aid int) bool {
	err := db.Driver.GetOne("account", aid, &r.Account)
	if err != nil {
		return false
	}
	return true
}
