package account

import (
	_ "encoding/json"
	"fmt"
	accountEnums "grpc-demo/enums/account"
	accountLogic "grpc-demo/logics/account"
	paramsUtils "grpc-demo/utils/params"

	//"fmt"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	AccountException "grpc-demo/exceptions/account"
	_ "grpc-demo/logics/resource"
	"grpc-demo/models/db"
	_ "net/http"
	_ "strings"
)

func GetAccountList(ctx iris.Context, auth authbase.AuthAuthorization) {

	auth.CheckAdmin()

	var lists []struct {
		Id         int   `json:"id"`
		UpdateTime int64 `json:"update_time"`
	}
	var count int
	table := db.Driver.Table("account")

	limit := ctx.URLParamIntDefault("limit", 10)
	page := ctx.URLParamIntDefault("page", 1)

	// 条件过滤
	if key := ctx.URLParam("key"); len(key) > 0 {
		keyString := fmt.Sprintf("%%%s%%", key)
		table = table.Where("nickname like ? or email like ?", keyString, keyString)
	}

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, update_time").Find(&lists)
	ctx.JSON(iris.Map{
		"accounts": lists,
		"total":    count,
		"limit":    limit,
		"page":     page,
	})
}

func PutAccount(ctx iris.Context, auth authbase.AuthAuthorization, aid int) {

	auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	logic := accountLogic.NewAccountLogic(auth, aid)
	account := logic.AccountModel()

	if err := db.Driver.GetOne("account", aid, &account); err != nil {
		panic(AccountException.AccountNotFount())
	}
	//if !auth.IsAdmin() && account.Id != auth.AccountModel().Id {
	//	panic(AccountException.NoPermission())
	//}
	params.Diff(account)

	if params.Has("role") {
		accountEnum := accountEnums.NewRoleEnums()
		if accountEnum.Has(params.Int("role", "角色")) {
			account.Role = int16(params.Int("role", "角色"))
		}
	}

	if params.Has("new_phone") && auth.AccountModel().Id == account.Id {
		newPhoneNum := params.Str("new_phone", "新号码")
		account.Phone = newPhoneNum
	}
	if params.Has("nick_name") && auth.AccountModel().Id == account.Id || auth.AccountModel().Role==1024{
		newNickname := params.Str("nick_name", "昵称")
		account.Nickname = newNickname
	}
	ctx.JSON(iris.Map{
		"id": account.Id,
	})
	db.Driver.Save(&account)

}

func MgetAccount(ctx iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")
	logic := accountLogic.NewAccountLogic(auth)
	data := make([]interface{}, 0)
	accounts := db.Driver.GetMany("account", ids, db.Account{})
	if auth.IsAdmin() == true{
		for _, account := range accounts {
			logic.SetAccountModel(account.(db.Account))
			func(data *[]interface{}) {
				*data = append(*data, logic.GetAccountInfo(1024))
				defer func() {
					recover()
				}()
			}(&data)
		}
	}else{
		for _, account := range accounts {
			logic.SetAccountModel(account.(db.Account))
			func(data *[]interface{}) {
				*data = append(*data, logic.GetAccountInfo(1))
				defer func() {
					recover()
				}()
			}(&data)
		}
	}




	ctx.JSON(data)
}
