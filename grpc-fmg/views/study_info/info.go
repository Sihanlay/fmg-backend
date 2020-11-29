package study_info

import (
	"github.com/Masterminds/squirrel"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	NewsException "grpc-demo/exceptions/study_info"
	"grpc-demo/models/db"
	logUtils "grpc-demo/utils/log"
	paramsUtils "grpc-demo/utils/params"
	qiniuUtils "grpc-demo/utils/qiniu"
)

func CreatStudyInfo(ctx iris.Context, auth authbase.AuthAuthorization) {
	auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var news db.News
	title := params.Str("title", "标题")
	content := params.Str("content", "内容")
	tagIds := params.List("tag_list","tag_list")
	kindTagMnt(tagIds,news)
	news = db.News{
		Title:           title,
		Content: content,
	}
	db.Driver.Create(news)
	ctx.JSON(iris.Map{
		"id": news.ID,
	})
}
//标签挂载
func kindTagMnt(k []interface{}, news db.News) {
	var tags []db.CommentTag
	if err := db.Driver.Where("id in (?)", k).Find(&tags).Error; err != nil || len(tags) == 0 {
		logUtils.Println(err)
		return
	}

	sql := squirrel.Insert("news_and_tag").Columns(
		"goods_id", "goods_tag_id", "tag_type",
	)

	for _, tag := range tags {
		sql = sql.Values(
			news.ID,
			tag.ID,
		)
	}
	if s, args, err := sql.ToSql(); err != nil {
		logUtils.Println(err)
	} else {
		if err := db.Driver.Exec(s, args...).Error; err != nil {
			logUtils.Println(err)
			return
		}
	}
}

func PutStudyInfo(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {

	auth.CheckLogin()
	var new db.News
	if err := db.Driver.Where("id = ?", cid).First(&new); err != nil {
		panic(accountException.AccountCarNotFount())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(new)
	new.Content = params.Str("content", "内容")
	new.Title = params.Str("title","title")

	db.Driver.Save(&new)
	ctx.JSON(iris.Map{
		"id": new.ID,
	})

}

func MgetStudyInfo(ctx iris.Context, auth authbase.AuthAuthorization) {

	auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids","ids")
	var news []db.News

	db.Driver.Where("id = ?", ids).Find(&news)
	data := make([]interface{}, 0, len(news))

	for _, new := range news {
		func(data *[]interface{}) {
			v := paramsUtils.ModelToDict(new, []string{"ID", "Title", "Content", "CreateTime"})
			*data = append(*data, v)
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}

func DeleteStudyInfo(ctx iris.Context, auth authbase.AuthAuthorization, nid int) {
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

func ListStudyInfos(ctx iris.Context, auth authbase.AuthAuthorization) {
	auth.CheckLogin()
	ctx.Text(qiniuUtils.GetUploadToken())

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
