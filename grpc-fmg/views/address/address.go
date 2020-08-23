package address

import (
	_ "context"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"

	//"encoding/json"
	//"errors"
	//"fmt"
	_ "grpc-demo/exceptions/system"
	_ "grpc-demo/models/db"

	_ "net/http"
	//"strings"
)

func GetAddress(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	addresses := make([]interface{}, 0)
	db.Driver.Where("AccountId = ?", uid).Find(&addresses)
	data := make([]interface{}, 0, len(addresses))
	for _, address := range addresses {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(address, []string{"ID", "ProvinceID", "CountryId", "CityID",
				"DistrictID", "Detail"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}

func ListAddress(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

}

func CreatAddress(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	province := params.Int("province_id", "province_id")
	country := params.Int("country_id", "country_id")
	city := params.Int("city_id", "city_id")
	district := params.Int("district_id", "district_id")
	detail := params.Str("detail", "detail")

	var address db.Address
	address = db.Address{
		CountryId:  country,
		ProvinceID: province,
		CityID:     city,
		DistrictID: district,
		Detail:     detail,
		AccountId:  uid,
	}

	db.Driver.Create(&address)
	ctx.JSON(iris.Map{
		"id": address.ID,
	})

}

func PutAddress(ctx iris.Context, auth authbase.AuthAuthorization, aid int) {
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	var address db.Address
	if err := db.Driver.GetOne("address", aid, &address); err != nil {
		panic(accountException.AddressNotFount())
	}
	if params.Has("city_id") {
		newCity := params.Int("city_id", "新地址")
		address.CityID = newCity
	}
	params.Diff(address)
	ctx.JSON(iris.Map{
		"id": address.ID,
	})
	db.Driver.Save(&address)
}
