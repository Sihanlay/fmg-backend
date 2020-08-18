package authbase

import (
	"fmt"
	"github.com/kataras/iris"
	"grpc-demo/constants"
	accountEnums "grpc-demo/enums/account"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	"grpc-demo/utils/hash"
	"time"
)

type clientAuthAuthorization struct {
	Account db.Account
	isLogin bool
	Context iris.Context
}

func (r *clientAuthAuthorization) AccountModel() *db.Account {
	return &r.Account
}

func (r *clientAuthAuthorization) CheckLogin() {
	if !r.isLogin {
		panic(accountException.AuthIsNotLogin())
	}
}

func (r *clientAuthAuthorization) IsLogin() bool {
	return r.isLogin
}

func (r *clientAuthAuthorization) IsAdmin() bool {
	return r.Account.Role == accountEnums.RoleAdmin
}

func (r *clientAuthAuthorization) CheckAdmin() {
	r.CheckLogin()
	if !r.IsAdmin() {
		panic(accountException.NoPermission())
	}
}

func (r *clientAuthAuthorization) LoadAuthenticationInfo() {
	r.loadFromHeader()
}

// 从header载入登录信息
func (r *clientAuthAuthorization) loadFromHeader() bool {
	defer func() {
		recover()
	}()
	token := r.Context.GetHeader(constants.ApiToken)
	fmt.Println(token)
	if token == "" {
		return false
	}
	var cookieStruct cookieInfo
	hash.DecodeToken(token, &cookieStruct)
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
func (r *clientAuthAuthorization) SetSession(aid int) {
	panic("client mode no session")
}

// 设置cookie
func (r *clientAuthAuthorization) SetCookie(aid int) string {
	payload := GenerateToken(aid, constants.CookieExpires)
	return payload
}

func (r *clientAuthAuthorization) fetchAccount(aid int) bool {
	err := db.Driver.GetOne("account", aid, &r.Account)
	if err != nil {
		return false
	}
	return true
}
