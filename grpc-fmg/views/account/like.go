package account

import (
	"fmt"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

func CreatLike(ctx iris.Context, auth authbase.AuthAuthorization,  gid int) {
	auth.CheckLogin()
	uid :=auth.AccountModel().Id
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var like db.Like

	goodsSpecificationId := params.Int("goods_specification_id", "商品规格id")


	like = db.Like{
		GoodsSpecificationId: goodsSpecificationId,
		AccountId:            uid,
		GoodsId:              gid,
	}
	db.Driver.Create(&like)
	ctx.JSON(iris.Map{
		"id": like.Id,
	})
}

func Mgetlike(ctx iris.Context, auth authbase.AuthAuthorization) {
	auth.CheckLogin()
	uid :=auth.AccountModel().Id
	type data struct {
		Ids []int `json:"ids"`
	}

	var likes []db.Like
	db.Driver.Where("account_id = ?", uid).Find(&likes)

	likedata := make([]interface{}, 0, len(likes))

	for _, like := range likes {

		func(likedata *[]interface{}) {
			info := paramsUtils.ModelToDict(like, []string{"Id", "GoodsId", "GoodsSpecificationId",
				"CreateTime", "GoodsSpecification"})

			*likedata = append(*likedata, info)
			defer func() {
				recover()
			}()
		}(&likedata)
	}


	ctx.JSON(iris.Map{
		"data":             likedata,

	})
}


func MDeleteLike(ctx iris.Context, auth authbase.AuthAuthorization) {
	auth.CheckLogin()

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids","收藏列表")
	if err := db.Driver.Debug().Exec("delete from like where id in (?)",ids).Error; err != nil {

		fmt.Println(err)
		panic(accountException.LikeNotFount())
	}

	ctx.JSON(iris.Map{
		"id": ids,
	})
}

