package account

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"grpc-demo/constants"
	authbase "grpc-demo/core/auth"
	"grpc-demo/core/cache"
	AccountException "grpc-demo/exceptions/account"
	resourceLogic "grpc-demo/logics/resource"
	"grpc-demo/models/db"
	"grpc-demo/utils"
	"grpc-demo/utils/hash"
	paramsUtils "grpc-demo/utils/params"
	"io/ioutil"
	"net/http"
	"strings"
	//"grpc-demo/wechat
)

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func Wxlogin(js_code string) (*WXLoginResp, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	Code2SessURL := "https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&js_code={code}&grant_type=authorization_code"
	Code2SessURL = strings.Replace(Code2SessURL, "{appid}", "wx1328c016e69fdf9f", -1)
	Code2SessURL = strings.Replace(Code2SessURL, "{secret}", "820d7c06cf3582fa4c7c55c4b559afb0", -1)
	Code2SessURL = strings.Replace(Code2SessURL, "{code}", js_code, -1)
	resp, err := client.Get(Code2SessURL)
	////关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	wxResp := WXLoginResp{}
	err = json.NewDecoder(resp.Body).Decode(&wxResp)
	if err != nil {
		return nil, err
	}

	return &wxResp, nil

}

func WxFirstLogin(c iris.Context, openid string) {

	v := hash.GetRandomString(10)

	_, _ = cache.Redis.Do(constants.DbNumberModel, "set", v, openid)

	c.JSON(iris.Map{
		"err_msg": AccountException.LoginFailed(),
		"key":     v,
	})

}

func Register(c iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(c))
	key := params.Str("key", "key")
	var account db.Account
	openid, err := redis.String(cache.Redis.Do(constants.DbNumberModel, "get", key))

	if params.Has("city") {
		city := params.Str("city", "city")
		account.City = city
	}
	if params.Has("country") {
		country := params.Str("country", "country")
		account.City = country
	}
	if params.Has("province") {
		province := params.Str("province", "province")
		account.City = province
	}
	if params.Has("nickName") {
		nickname := params.Str("nickName", "nickName")
		account.Nickname = nickname
	}
	if err == nil && openid != "" {
		account = db.Account{
			OpenId: openid,
		}
		db.Driver.Create(&account)
	} else {
		panic(AccountException.AccountNotFount())
	}
	token := auth.SetCookie(account.Id)
	c.JSON(iris.Map{
		"id":     account.Id,
		"token":  token,
		"openid": openid,
	})

}

func Login(c iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(c))
	//c.Text(qiniuUtils.GetUploadToken())
	//判断小程序登录
	mode := c.GetHeader(constants.ApiMode)

	if mode == "client" {
		code := params.Str("js_code", "js_code") //  获取code
		wxLoginResp, err := Wxlogin(code)
		fmt.Println(wxLoginResp.ErrCode)
		fmt.Println(wxLoginResp.ErrMsg)
		if wxLoginResp.OpenId == "" {
			panic(AccountException.LoginFailed())
		}
		fmt.Println(err)
		if err != nil {
			panic(AccountException.AccountNotFount())
			return
		}

		// 根据code获取 openID 和 session_key

		//判断openid是否存在数据库
		// 没有就创建model保存登录态
		var account db.Account

		if err := db.Driver.
			Where("open_id = ?", wxLoginResp.OpenId).
			First(&account).Error; err != nil {

			WxFirstLogin(c, wxLoginResp.OpenId)

		} else {
			token := auth.SetCookie(account.Id)
			auth.IsLogin()
			c.JSON(iris.Map{
				"id":    account.Id,
				"token": token,

				"open_id": wxLoginResp.OpenId,
			})
		}
	} else {
		//web端登录
		panic(AccountException.AccountNotFount())

	}

}

//web登录
func WebLogin(c iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(c))
	var account db.Account
	openid := params.Str("open_id", "open_id")
	if err := db.Driver.
		Where("open_id = ?", openid).
		First(&account).Error; err != nil {

		account = db.Account{
			OpenId: openid,
		}
		db.Driver.Create(&account)
		auth.SetCookie(account.Id)

	} else {
		auth.SetCookie(account.Id)
	}
	c.JSON(iris.Map{
		"id": account.Id,
	})
}

// 获取头像数据
func getAvator(tx *gorm.DB, url string, account *db.Account) bool {
	if response, err := utils.Requests("GET", url, nil); err == nil && response.StatusCode == http.StatusOK {
		if body, err := ioutil.ReadAll(response.Body); err == nil {
			defer response.Body.Close()
			logic := resourceLogic.NewReousrcesLocalStorage("account_avator")
			account.Avator = logic.SaveFile(fmt.Sprintf("%d/%s", account.Id, "avator.jpg"), body, true)
		}
		if err := tx.Save(&account).Error; err != nil {
			return false
		}
	}
	return true
}

func Logout(ctx iris.Context, auth authbase.AuthAuthorization) {
	auth.SetCookie(0)
	ctx.JSON(iris.Map{
		"status": "success",
	})
}
