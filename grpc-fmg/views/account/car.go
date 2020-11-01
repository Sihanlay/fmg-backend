package account

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
	requestsUtils "grpc-demo/utils/requests"
)

func CreatAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, uid int, gid int) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var car db.AccountCar
	goodsCount := params.Int("goods_count", "商品数量")
	goodsSpecificationId := params.Int("goods_specification_id", "商品规格id")
	deliveryKind := params.Int("delivery_kind", "发货方式")

	car = db.AccountCar{
		GoodsCount:           goodsCount,
		GoodsSpecificationId: goodsSpecificationId,
		AccountId:            uid,
		GoodsId:              gid,
		DeliveryKind:         deliveryKind,
	}
	db.Driver.Create(&car)
	ctx.JSON(iris.Map{
		"id": car.Id,
	})
}

func MgetAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	type data struct {
		Ids []int `json:"ids"`
	}

	type goodsInfo struct {
		GoodId          int
		SpecificationId int
		Count           int
		Delivery        int
	}
	var cars []db.AccountCar
	db.Driver.Where("account_id = ?", uid).Find(&cars)

	var goodsInfoList []goodsInfo
	goodsIds := make([]int, 0)
	cardata := make([]interface{}, 0, len(cars))
	var count int = 0
	for _, car := range cars {
		if car.IsCheck == true {
			goodsIds = append(goodsIds, car.GoodsId)
			goodsInfoList = append(goodsInfoList, goodsInfo{car.GoodsId, car.GoodsSpecificationId, car.GoodsCount, car.DeliveryKind})
			count += car.GoodsCount
		}
		func(cardata *[]interface{}) {
			info := paramsUtils.ModelToDict(car, []string{"Id", "GoodsId", "GoodsCount", "GoodsSpecificationId",
				"CreateTime", "GoodsSpecification", "DeliveryKind",
				"IsCheck"})

			*cardata = append(*cardata, info)
			defer func() {
				recover()
			}()
		}(&cardata)
	}

	fmt.Println(goodsIds)
	newdata := data{}
	newdata.Ids = goodsIds
	fmt.Println(newdata)
	reqdata, _ := json.Marshal(newdata)
	fmt.Println(string(reqdata))
	v, err := requestsUtils.Do("POST", "/goods/_mget", reqdata)
	if err != nil {
		fmt.Print(err)
		return
	}
	var sumPrice float32 = 0
	var totalDiscountSum float32 = 0
	var maxDelivery float32 = 0
	var goodsPrice float32 = 0
	var discountSum float32 = 0
	for i, item := range v.([]interface{}) {
		//fmt.Println(goodsInfoList[i].SpecificationId)
		if float32(item.(map[string]interface{})["carriage"].(float64)) >= maxDelivery {
			maxDelivery = float32(item.(map[string]interface{})["carriage"].(float64))
		}

		for _, gitem := range item.(map[string]interface{})["specification"].([]interface{}) {
			if int(gitem.(map[string]interface{})["id"].(float64)) == goodsInfoList[i].SpecificationId {
				sumPrice += float32(gitem.(map[string]interface{})["price"].(float64) * (float64(goodsInfoList[i].Count)))
				goodsPrice += float32(gitem.(map[string]interface{})["price"].(float64) * (float64(goodsInfoList[i].Count)))
				if (item.(map[string]interface{})["sale"].(bool)) == true {
					totalDiscountSum += float32(gitem.(map[string]interface{})["reduced_price"].(float64) * (float64(goodsInfoList[i].Count)))
				} else {
					totalDiscountSum += float32(gitem.(map[string]interface{})["price"].(float64) * (float64(goodsInfoList[i].Count)))
				}

			}
		}
		sumPrice += maxDelivery
		discountSum = totalDiscountSum
		totalDiscountSum += maxDelivery

	}

	ctx.JSON(iris.Map{
		"data":             cardata,
		"count":            count,
		"sum":              sumPrice,
		"totalDiscountSum": totalDiscountSum,
		"goodsPrice":       goodsPrice,
		"discountSum":      discountSum,
	})
}

func MsetAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	//路由加账户id
	//判断购物车id是否属于此账户
	var ids []struct {
		Id int `json:"id"`
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	//是否全选
	if params.Has("flag") {
		db.Driver.Table("account_car").Debug().Select("id").Where("account_id = ?", uid).Find(&ids)
		check_ids := make([]int, 0)

		for _, id := range ids {
			check_ids = append(check_ids, id.Id)
		}

		check := params.Bool("flag", "flag")
		db.Driver.Table("account_car").Debug().Where("id IN (?)", check_ids).Updates(map[string]interface{}{"is_check": check})

	}

	ctx.JSON(iris.Map{
		"status": "success",
	})

}

//func MsetAccountCar(ctx iris.Context, auth authbase.AuthAuthorization,uid int){
//
//	//路由加账户id
//	//判断购物车id是否属于此账户
//	var ids []struct {
//		Id         int   `json:"id"`
//	}
//
//	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
//
//	//是否全选
//	if params.Has("flag"){
//		db.Driver.Table("account_car").Debug().Select("id").Where("account_id = ?", uid).Find(&ids)
//		check_ids := make([]int,0)
//
//		for _, id := range ids {
//			check_ids = append(check_ids, id.Id)
//		}
//
//		check := params.Bool("flag","flag")
//		db.Driver.Table("account_car").Debug().Where("id IN (?)", check_ids).Updates(map[string]interface{}{"is_check":check})
//
//	}else if params.Has("info"){
//		info := params.List("info","info")
//		for _, data := range info {
//			var car db.AccountCar
//			cid := int(data.(map[string]interface{})["id"].(float64))
//			if err := db.Driver.GetOne("account_car",cid,&car); err != nil {
//				panic(accountException.AccountCarNotFount())
//			}
//			params.Diff(car)
//			car.GoodsCount = params.Int("goods_count","商品数量")
//			car.GoodsSpecificationId = params.Int("goods_specification_id","商品规格id")
//			car.IsCheck = params.Bool("is_check","is_check")
//			car.DeliveryKind = params.Int("delivery_kind","发货方式")
//
//			db.Driver.Save(&car)
//
//		}
//	}
//
//	ctx.JSON(iris.Map{
//		"status": "success",
//	})
//
//}
func PutAccountCar(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var car db.AccountCar
	if err := db.Driver.GetOne("account_car", cid, &car); err != nil {
		panic(accountException.AccountCarNotFount())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(car)

	//car.GoodsName = params.Str("goods_name","goods_name")
	car.GoodsCount = params.Int("goods_count", "商品数量")
	//car.GoodsPrice = params.Int("goods_price","goods_price")
	//car.Picture = params.Str("goods_pictures","商品图片")
	car.GoodsSpecificationId = params.Int("goods_specification_id", "商品规格id")
	car.IsCheck = params.Bool("check", "check")
	car.DeliveryKind = params.Int("delivery_kind", "发货方式")

	db.Driver.Save(&car)
	ctx.JSON(iris.Map{
		"id": car.Id,
	})

}

func DeleteGoodsCar(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var car db.AccountCar
	if err := db.Driver.GetOne("accountcar", cid, &car); err == nil {

		db.Driver.Delete(car)
	} else {
		panic(accountException.AccountCarNotFount())
	}

	ctx.JSON(iris.Map{
		"id": cid,
	})
}
