package deliever

import (
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

func CreatDelivery(ctx iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	order_code := params.Int("order_code", "order_code")
	delivery_corp_name := params.Str("delivry_corp_name", "delivery_corp_name")
	delivery_sheet_code := params.Int("delivry_sheet_code", "delivery_sheet_code")

	invoice_status := params.Int("invoice_status", "invoice_status")

	var delievery db.Delivery
	delievery = db.Delivery{
		OrderCode: order_code,
		DeliveryCorpName: delivery_corp_name,
		DeliverySheetCode: delivery_sheet_code,
		InvoiceStatus: invoice_status,
	}

	db.Driver.Create(&delievery)
	ctx.JSON(iris.Map{
		"id": delievery.ID,
	})

}

func MgetDelivery(ctx iris.Context, auth authbase.AuthAuthorization) {


	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	ids := params.List("ids", "id列表")
	deliveries := db.Driver.GetMany("delivery", ids,db.Delivery{})
	data := make([]interface{}, 0,len(ids))
	for _, delivery := range deliveries {

		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(delivery,[]string{
				"Id","OrderCode","DeliveryCorpName","DeliverySheetCode","InvoiceStatus:",
			}))
				defer func() {
				recover()

			}()
		}(&data)
	}
	ctx.JSON(data)
}
