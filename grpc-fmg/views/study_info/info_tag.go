package study_info

import (
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	NewsException "grpc-demo/exceptions/study_info"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

func CreatStudyInfoTag(ctx iris.Context, auth authbase.AuthAuthorization) {
	//auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var newsTag db.NewsTag
	name := params.Str("name", "标题")

	newsTag = db.NewsTag{
		Name:           name,

	}
	db.Driver.Create(&newsTag)
	ctx.JSON(iris.Map{
		"id": newsTag.ID,
	})
}
func MgetStudyInfoTag(ctx iris.Context, auth authbase.AuthAuthorization) {

	//auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids","ids")

	news := db.Driver.GetMany("news_tag", ids,db.NewsTag{})
	data := make([]interface{}, 0, len(ids))

	for _, item := range news {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(item, []string{"ID", "Title", "Content", "CreateTime"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}

func DeleteStudyInfoTag(ctx iris.Context, auth authbase.AuthAuthorization, nid int) {
	var new db.News
	if err := db.Driver.GetOne("news", nid, &new); err == nil {

		db.Driver.Delete(new)
	}else{
		panic(NewsException.InfoNotExists())
	}

	ctx.JSON(iris.Map{
		"id": nid,
	})
}

func ListStudyInfosTag(ctx iris.Context, auth authbase.AuthAuthorization) {
	//auth.CheckLogin()
	//ctx.Text(qiniuUtils.GetUploadToken())

	var lists []struct {
		Id         int   `json:"id"`
		UpdateTime int64 `json:"update_time"`
	}
	var count int

	table := db.Driver.Table("news")

	limit := ctx.URLParamIntDefault("limit", 10)
	page := ctx.URLParamIntDefault("page", 1)


	//销售标签过滤
	//if saleTag := ctx.URLParamIntDefault("sale_tag", 0); saleTag != 0 {
	//	var tag db.SaleTag
	//	if err := db.Driver.GetOne("sale_tag", saleTag, &tag); err == nil {
	//		var goods []db.GoodsAndTag
	//
	//		if err := db.Driver.Where("goods_tag_id = ? and tag_type = ?", saleTag, goodsTagEnum.TagTypeSaleTag).Find(&goods).Error; err == nil {
	//			ids := make([]interface{}, len(goods))
	//			for index, v := range goods {
	//				ids[index] = v.GoodsID
	//			}
	//			table = table.Where("id in (?)", ids)
	//		}
	//	}
	//}

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, update_time").Find(&lists)
	ctx.JSON(iris.Map{
		"news": lists,
		"total": count,
		"limit": limit,
		"page":  page,
	})

}

