package comment

import (
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

func GetComment(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var comment db.Comment
	db.Driver.Where("id = ?", cid).First(&comment)

	info := paramsUtils.ModelToDict(comment, []string{"ID", "OrderID", "Content"})

	ctx.JSON(info)

}
func CreatComment(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var comment db.Comment
	orderId := params.Int("order_id", "订单id")
	content := params.Str("content", "内容")

	comment = db.Comment{

		AuthorID: uid,
		OrderID:  orderId,
		Content:  content,
	}
	db.Driver.Create(&comment)
	ctx.JSON(iris.Map{
		"id": comment.ID,
	})
}

func MgetComment(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	var comments []db.Comment

	db.Driver.Where("account_id = ?", uid).Find(&comments)
	data := make([]interface{}, 0, len(comments))

	for _, comment := range comments {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(comment, []string{"ID", "OrderID", "AuthorID", "Content"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}

func PutComment(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var comment db.Comment
	if err := db.Driver.Where("id = ?", cid).First(&comment); err != nil {
		panic(accountException.AccountCarNotFount())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(comment)

	comment.Content = params.Str("comment", "评论")

	db.Driver.Save(&comment)
	ctx.JSON(iris.Map{
		"id": comment.ID,
	})

}

func DeleteComment(ctx iris.Context, auth authbase.AuthAuthorization, cid int) {
	var comment db.Comment
	if err := db.Driver.GetOne("accountcar", cid, &comment); err == nil {

		db.Driver.Delete(comment)
	} else {
		panic(accountException.AccountCarNotFount())
	}

	ctx.JSON(iris.Map{
		"id": cid,
	})
}
