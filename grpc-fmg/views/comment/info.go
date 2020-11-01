package comment

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
	accountException "grpc-demo/exceptions/account"
	"grpc-demo/models/db"
	paramsUtils "grpc-demo/utils/params"
)

func GetComment(ctx iris.Context, auth authbase.AuthAuthorization, oid int) {
	var comment db.Comment
	db.Driver.Where("order_id = ?", oid).First(&comment)

	info := paramsUtils.ModelToDict(comment, []string{"ID", "OrderID", "Content", "CommentTag", "Pictures"})

	ctx.JSON(info)

}
func CreatComment(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {
	//auth.CheckLogin()
	//order
	//account_id
	//auth.AccountModel().Id

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var comment db.Comment
	orderId := params.Int("order_id", "订单id")
	content := params.Str("content", "内容")
	tag := params.Int("tag", "表情")

	tx := db.Driver.Begin()

	comment = db.Comment{
		AuthorID:   uid,
		OrderID:    orderId,
		Content:    content,
		CommentTag: tag,
	}
	tx.Create(&comment)

	if params.Has("pictures") {
		pictures := params.List("pictures", "图片")
		if p, err := json.Marshal(pictures); err != nil {
			panic("序列化失败")
		} else {
			comment.Pictures = string(p)
			fmt.Println(string(p))
			if err := tx.Save(&comment).Debug().Error; err != nil {
				fmt.Println(err)
				panic("保存失败")
			}
		}
	}

	tx.Commit()

	ctx.JSON(iris.Map{
		"id": comment.ID,
	})
}

func MgetComment(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

	var comments []db.Comment

	db.Driver.Where("author_id = ?", uid).Find(&comments)
	data := make([]interface{}, 0, len(comments))

	for _, comment := range comments {
		func(data *[]interface{}) {
			v := paramsUtils.ModelToDict(comment, []string{"ID", "OrderID", "AuthorID", "Content", "CommentTag"})
			var p []interface{}
			if err := json.Unmarshal([]byte(comment.Pictures), &p); err != nil {
				fmt.Println(p)
				panic("反序列化失败")
			}
			v["pictures"] = p
			*data = append(*data, v)
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
	if params.Has("pictures") {
		pictures := params.List("pictures", "图片")
		v, _ := json.Marshal(pictures)
		comment.Pictures = string(v)
	}

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
