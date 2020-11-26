package accountLogic

import (
	authbase "grpc-demo/core/auth"
	"grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

var field = []string{
	"Nickname", "Email", "Id", "Role", "Phone", "PhoneValidated", "UpdateTime",
	"EmailValidated", "Avator", "Motto", "CreateTime", "Init",
}

var normalfield = []string{
	"Nickname","Id", "Avator", "Motto","CreateTime",
}

type AccountLogic struct {
	auth    authbase.AuthAuthorization
	account db.Account
}

func NewAccountLogic(auth authbase.AuthAuthorization, aid ...int) AccountLogic {
	var account db.Account

	if len(aid) > 0 {
		if err := db.Driver.GetOne("account", aid[0], &account); err != nil {
			panic(accountException.AccountNotFount())
		}
	} else {
		account = *auth.AccountModel()
	}
	return AccountLogic{
		account: account,
		auth:    auth,
	}
}

func (a *AccountLogic) SetAccountModel(account db.Account) {
	a.account = account
}

func (a *AccountLogic) AccountModel() *db.Account {
	return &a.account
}

func (a *AccountLogic) GetAccountInfo(role int) interface{} {

	//if len(a.account.Avator) > 0 {
	//	a.account.Avator = resourceLogic.GenerateToken(a.account.Avator, -1, -1)
	//}
	var data map[string]interface{}
	if role == 1024{
		data = paramsUtils.ModelToDict(a.account, field)
	}else{
		data = paramsUtils.ModelToDict(a.account, normalfield)
	}

	//data["oauth"] = a.GetOauth()
	return data
}

//// 获取已验证的第三方
//func (a *AccountLogic) GetOauth() map[int]string {
//
//	oauth := make(map[int]string)
//	if rows, err := db.Driver.
//		Table("account_oauth").
//		Where("account_id = ?", a.account.ID).
//		Select("model, user_info").Rows(); err == nil {
//		for rows.Next() {
//			var model int
//			var userInfo string
//			if err := rows.Scan(&model, &userInfo); err == nil {
//				var user interface{}
//				if err := json.Unmarshal([]byte(userInfo), &user); err == nil {
//					switch model {
//					case accountEnums.OauthWejudge:
//						oauth[model] = user.(map[string]interface{})["email"].(string)
//					}
//				}
//			}
//		}
//	}
//	return oauth
//}

