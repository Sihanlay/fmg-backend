package authbase

import (
	"fmt"
	"github.com/kataras/iris"
	"grpc-demo/constants"
	"grpc-demo/models/db"
	"grpc-demo/utils/hash"
	"time"
)

type cookieInfo struct {
	AccountId  int   `json:"account_id"`
	ExpireTime int64 `json:"expire_time"`
}

type AuthAuthorization interface {
	SetSession(aid int)
	SetCookie(aid int) string
	CheckLogin()
	IsAdmin() bool
	CheckAdmin()
	AccountModel() *db.Account
	IsLogin() bool
	LoadAuthenticationInfo()
}

func NewAuthAuthorization(ctx iris.Context) AuthAuthorization {
	mode := ctx.GetHeader(constants.ApiMode)

	var authorization AuthAuthorization
	fmt.Println(mode)
	if mode == "client" {
		authorization = &clientAuthAuthorization{
			isLogin: false,
			Context: ctx,
		}
	} else {
		authorization = &dAuthAuthorization{
			isLogin: false,
			Context: ctx,
		}
	}
	authorization.LoadAuthenticationInfo()
	return authorization
}

// 生成token
func GenerateToken(aid int, expire int64) string {
	payload := cookieInfo{
		AccountId:  aid,
		ExpireTime: expire + time.Now().Unix(),
	}
	return hash.GenerateToken(payload, true)
}
