package account

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"grpc-demo/constants"
	authbase "grpc-demo/core/auth"
	AccountException "grpc-demo/exceptions/account"
	resourceLogic "grpc-demo/logics/resource"
	"grpc-demo/models/db"
	"grpc-demo/utils"
	paramsUtils "grpc-demo/utils/params"
	qiniuUtils "grpc-demo/utils/qiniu"
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

	Code2SessURL := "https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&js_code={code}&grant_type=authorization_code"
	Code2SessURL = strings.Replace(Code2SessURL, "{appid}", "wx1328c016e69fdf9f", -1)
	Code2SessURL = strings.Replace(Code2SessURL, "{secret}", "498352d438f07243860b8dd54ef946f0", -1)
	Code2SessURL = strings.Replace(Code2SessURL, "{code}", js_code, -1)
	resp, err := http.Get(Code2SessURL)
	////关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	wxResp := WXLoginResp{}
	err = json.NewDecoder(resp.Body).Decode(&wxResp)
	if err != nil {
		return nil, err
	}

	return &wxResp, nil

}

func Login(c iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(c))
	c.Text(qiniuUtils.GetUploadToken())
	//判断小程序登录
	mode := c.GetHeader(constants.ApiMode)

	if mode == "client" {
		code := params.Str("js_code", "js_code") //  获取code
		// 根据code获取 openID 和 session_key
		wxLoginResp, err := Wxlogin(code)
		if err != nil {
			c.JSON(AccountException.AccountNotFount())
			return
		}
		//判断openid是否存在数据库
		// 没有就创建model保存登录态
		var account db.Account

		if err := db.Driver.
			Where("open_id = ?", wxLoginResp.OpenId).
			First(&account).Error; err != nil {

			//nickname := params.Str("nickname","昵称")
			account = db.Account{
				OpenId: wxLoginResp.OpenId,
				//Nickname: nickname,
			}
		}
		token := auth.SetCookie(account.Id)
		c.JSON(iris.Map{
			"id":    account.Id,
			"token": token,
		})
	} else {
		//web端登录
		var account db.Account
		db.Driver.Where(
			"email = ? and email_validated = true and password = ?",
			params.Str("email", "邮箱"),
			params.Str("password", "密码"),
		).First(&account)

		auth.SetCookie(account.Id)
	}

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
