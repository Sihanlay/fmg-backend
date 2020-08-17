package wechat

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"
)

func requestToken(appid, secret string) (string, error) {
	u, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	paras.Set("appid", appid)
	paras.Set("secret", secret)
	paras.Set("grant_type", "client_credential")
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", errors.New("request token err :" + err.Error())
	}

	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return "", errors.New("request token response json parse err :" + err.Error())
	}
	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
		accessToken, _ := jMap["access_token"].(string)
		return accessToken, nil
	} else {
		//返回错误信息
		errcode := jMap["errcode"].(string)
		errmsg := jMap["errmsg"].(string)
		err = errors.New(errcode + ":" + errmsg)
		return "", err
	}

}


func init() {
	appid := "wx9f8f8d7fc02ac384"
	secret := "a998eda987fe7b9c45fa93d5542e1161"
	freshTokenTicker := time.NewTicker(7000 * time.Second)
	//requestToken()

	go func() {

		for range freshTokenTicker.C {

			accessToken, err := requestToken(appid, secret)
			if err != nil {
				//TODO 错误处理
			}
			log.Printf("token refresh :%s", accessToken)
		}
	}()

}
