package account

import (
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

func CreatAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, uid int,gid int){


	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var car db.AccountCar
	//goodsName := params.Str("goods_name","goods_name")
	goodsCount := params.Int("goods_count","商品数量")
	//goodsPrice := params.Int("goods_price","goods_price")
	//goodsPictures := params.Str("goods_pictures","商品图片")
	goodsSpecificationId := params.Int("goods_specification_id","商品规格id")
	goodsSpecification := params.Str("goods_specification","商品规格")

	//if params.Has("goods_pictures"){
	//	car.Picture = params.Str("goods_pictures","商品图片")
	//}

	car = db.AccountCar{
		//GoodsName:          goodsName,
		GoodsCount:         goodsCount,
		GoodsSpecification: goodsSpecification,
		GoodsSpecificationId: goodsSpecificationId,
		//GoodsPrice: goodsPrice,
		AccountId:          uid,
		GoodsId:            gid,
		//Picture: goodsPictures,
	}
	db.Driver.Create(&car)
	ctx.JSON(iris.Map{
		"id": car.Id,
	})
}

func MgetAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, uid int){

	var cars []db.AccountCar

	db.Driver.Where("account_id = ?", uid).Find(&cars)
	data := make([]interface{}, 0, len(cars))

	for _, car := range cars {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(car, []string{"Id", "GoodsId", "GoodsCount","GoodsSpecificationId",
				"CreateTime","GoodsSpecification",
				"IsCheck"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}

func PutAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var car db.AccountCar
	if err := db.Driver.GetOne("account_car", cid, &car); err != nil {
		panic(accountException.AccountCarNotFount())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(car)

	//car.GoodsName = params.Str("goods_name","goods_name")
	car.GoodsCount = params.Int("goods_count","商品数量")
	//car.GoodsPrice = params.Int("goods_price","goods_price")
	//car.Picture = params.Str("goods_pictures","商品图片")
	car.GoodsSpecificationId = params.Int("goods_specification_id","商品规格id")
	car.IsCheck = params.Bool("check","check")

	db.Driver.Save(&car)
	ctx.JSON(iris.Map{
		"id": car.Id,
	})




}


func DeleteGoodsCar(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var car db.AccountCar
	if err := db.Driver.GetOne("accountcar", cid, &car); err == nil {

		db.Driver.Delete(car)
	}else{
		panic(accountException.AccountCarNotFount())
	}

	ctx.JSON(iris.Map{
		"id": cid,
	})
}