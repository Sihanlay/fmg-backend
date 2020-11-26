package deliever

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"strings"

	//"encoding/hex"
	"encoding/json"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	deliveryExceptions "grpc-demo/exceptions/delievery"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
	"io"
	"net/http"
)

func CreatDelivery(ctx iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	orderCode := params.Str("order_code", "order_code")
	deliveryCorpName := params.Str("delivry_corp_name", "delivery_corp_name")
	deliverySheetCode := params.Str("delivry_sheet_code", "delivery_sheet_code")
	invoiceStatus := params.Int("invoice_status", "invoice_status")
	creatTime := params.Int("create_time", "createTime")

	var order db.TestChildOrder
	if err := db.Driver.Where("order_num=?", orderCode).Find(&order).Error;err != nil{
		fmt.Println(err)
		panic(deliveryExceptions.OrderNotExists())
	}
	params.Diff(order)
	order.DeliveryTime = int64(creatTime)
	order.TrackingID = deliverySheetCode
	order.TrackingCompany = deliveryCorpName
	db.Driver.Save(&order)

	var delievery db.Delivery
	delievery = db.Delivery{
		OrderCode:         orderCode,
		DeliveryCorpName:  deliveryCorpName,
		DeliverySheetCode: deliverySheetCode,
		InvoiceStatus:     invoiceStatus,
	}

	db.Driver.Create(&delievery)
	ctx.JSON(iris.Map{
		"id": delievery.ID,
	})

}
func GetDeliveryList(ctx iris.Context, auth authbase.AuthAuthorization) {

	var lists []struct {
		Id        int    `json:"id"`
		OrderCode string `json:"order_code"`
	}
	var count int
	table := db.Driver.Table("delivery")

	limit := ctx.URLParamIntDefault("limit", 10)
	page := ctx.URLParamIntDefault("page", 1)

	//// 条件过滤
	//if key := ctx.URLParam("key"); len(key) > 0 {
	//	keyString := fmt.Sprintf("%%%s%%", key)
	//	table = table.Where("nickname like ? or email like ?", keyString, keyString)
	//}

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id,order_code").Find(&lists)
	ctx.JSON(iris.Map{
		"delivery": lists,
		"total":    count,
		"limit":    limit,
		"page":     page,
	})
}

func MgetDelivery(ctx iris.Context, auth authbase.AuthAuthorization) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")
	deliveries := db.Driver.GetMany("delivery", ids, db.Delivery{})
	data := make([]interface{}, 0, len(ids))
	for _, delivery := range deliveries {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(delivery, []string{
				"Id", "OrderCode", "DeliveryCorpName", "DeliverySheetCode", "InvoiceStatus:",
			}))
			defer func() {
				recover()

			}()
		}(&data)
	}
	ctx.JSON(data)
}

type paramData struct {
	Com      string `json:"com"`
	Num      string `json:"num"`
	From     string `json:"from"`
	Phone    string `json:"phone"`
	To       string `json:"to"`
	Resultv2 string `json:"resultv2"`
	Show     string `json:"show"`
	Order    string `json:"order"`
}

type respBody struct {
	Message   string        `json:"message"`
	State     string        `json:"state"`
	Status    string        `json:"status"`
	Condition string        `json:"condition"`
	Ischeck   string        `json:"ischeck"`
	Com       string        `json:"com"`
	Nu        string        `json:"nu"`
	Data      []interface{} `json:"data"`
}


const (
	customer = "4E59BDEDD9D6279BCE70939B5989DF54"
	key      = "flHKGrTQ6370"
)

func DeliveryInfo(ctx iris.Context, auth authbase.AuthAuthorization) error {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	req := make(map[string]interface{})
	deliveryCorpName := params.Str("delivry_corp_name", "delivery_corp_name")
	deliverySheetCode := params.Str("delivry_sheet_code", "delivery_sheet_code")

	param := paramData{}
	param.Com = deliveryCorpName
	param.Num = deliverySheetCode
	param.Order = "desc"
	param.Resultv2 = "0"
	param.Show = "0"
	req["customer"] = customer
	req["param"] = param

	//MD5加密

	h := md5.New()
	var buf bytes.Buffer
	info, _ := json.Marshal(param)
	buf.WriteString(string(info))
	buf.WriteString(key)
	buf.WriteString(customer)
	h.Write([]byte(buf.String()))
	keyInfo := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	req["sign"] = keyInfo
	//data, _ := json.Marshal(req)
	reqdata := fmt.Sprintf("customer=%s&sign=%s&param=%s", customer, keyInfo, string(info))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var reader io.Reader
	if reqdata != "" {
		reader = strings.NewReader(reqdata)
	}
	client := &http.Client{Transport: tr}
	apiresp := respBody{}
	resp, err := client.Post("https://poll.kuaidi100.com/poll/query.do", "application/x-www-form-urlencoded", reader)
	if resp != nil && resp.Body != nil {

		err = json.NewDecoder(resp.Body).Decode(&apiresp)
		details := apiresp.Data
		idata := make([]interface{}, 0, len(details))
		for _, item := range details {
			context := item.(map[string]interface{})["context"]
			time := item.(map[string]interface{})["time"]
			ftime := item.(map[string]interface{})["ftime"]
			func(idata *[]interface{}) {
				//info := data{}
				//info.Context = context
				//info.Time = time
				//info.Ftime = ftime
				*idata = append(*idata, context, time, ftime)
				defer func() {
					recover()
				}()
			}(&idata)

		}

		ctx.JSON(iris.Map{
			"info": apiresp,
		})
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}
