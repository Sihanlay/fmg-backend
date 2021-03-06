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

func MGetCommentByGood(ctx iris.Context, auth authbase.AuthAuthorization,gid int) {
	//auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var comments []db.Comment

	data := make([]interface{}, 0, len(comments))
	if params.Has("tag"){
		tag := params.Int("tag","tag")
		db.Driver.Where("good_id = ? and comment_tag = ?",gid,tag).Order("create_time desc").Find(&comments)
	}else{
		db.Driver.Where("good_id = ?",gid).Order("create_time desc").Find(&comments)
	}


	for _, comment := range comments {
		func(data *[]interface{}) {
			v := paramsUtils.ModelToDict(comment, []string{"ID", "GoodID", "AuthorID", "Content", "CommentTag","CreateTime"})
			var p []interface{}
			if comment.Pictures != ""{
				if err := json.Unmarshal([]byte(comment.Pictures), &p); err != nil {
					fmt.Println(p)
					panic("反序列化失败")
				}
				v["pictures"] = p
			}

			*data = append(*data, v)
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}
func MGetCommentByUser(ctx iris.Context, auth authbase.AuthAuthorization) {
	uid := auth.AccountModel().Id
	//auth.CheckLogin()
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var comments []db.Comment

	data := make([]interface{}, 0, len(comments))
	if params.Has("tag"){
		tag := params.Int("tag","tag")
		db.Driver.Where("account_id = ? and comment_tag = ?",uid,tag).Order("create_time desc").Find(&comments)
	}else{
		db.Driver.Where("account_id = ?",uid).Order("create_time desc").Find(&comments)
	}


	for _, comment := range comments {
		func(data *[]interface{}) {
			v := paramsUtils.ModelToDict(comment, []string{"ID", "GoodID", "AuthorID", "Content", "CommentTag","CreateTime"})
			var p []interface{}
			if comment.Pictures != ""{
				if err := json.Unmarshal([]byte(comment.Pictures), &p); err != nil {
					fmt.Println(p)
					panic("反序列化失败")
				}
				v["pictures"] = p
			}

			*data = append(*data, v)
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)

}
func CreatComment(ctx iris.Context, auth authbase.AuthAuthorization, gid int,oid int) {
	auth.CheckLogin()
	accountId := auth.AccountModel().Id

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	var comment db.Comment
	var orderDetail db.TestOrderDetail
	var order db.TestOrder

	err1 :=db.Driver.GetOne("test_order_detail",oid,&orderDetail);if err1 != nil{
		panic("找不到该订单")
	}
	err2 := db.Driver.GetOne("test_order_detail", orderDetail.OrderID, &order);if err2 != nil{
		panic("找不到该订单")
	}
	if order.Status == 3{
		panic("订单已完成，无法评论")
	}
	orderDetail.IsComment = 1
	db.Driver.Save(&order)


	Content := params.Str("content", "内容")
	comment.Content = Content

	tag := params.Int("tag", "标签")

	tx := db.Driver.Begin()

	comment = db.Comment{
		AuthorID:   accountId,
		GoodID:    gid,
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

func MgetComment(ctx iris.Context, auth authbase.AuthAuthorization) {

	auth.CheckLogin()
	uid := auth.AccountModel().Id
	var comments []db.Comment

	db.Driver.Where("author_id = ?", uid).Find(&comments)
	data := make([]interface{}, 0, len(comments))

	for _, comment := range comments {
		func(data *[]interface{}) {
			v := paramsUtils.ModelToDict(comment, []string{"ID", "GoodID", "AuthorID", "Content", "CommentTag"})
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

func PutComment(ctx iris.Context, auth authbase.AuthAuthorization, cid int,oid int) {
	auth.CheckLogin()
	var comment db.Comment
	if err := db.Driver.Where("id = ?", cid).First(&comment); err != nil {
		panic(accountException.AccountCarNotFount())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(comment)
	if params.Has("second_content") {
		secondContent := params.Str("second_content", "内容")
		comment.SecondContent = secondContent
		var order db.TestOrderDetail

		err :=db.Driver.GetOne("test_order_detail",oid,&order);if err != nil{
			panic("找不到该订单")
		}

		order.IsComment = 2
		db.Driver.Save(&order)
	}
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
	if err := db.Driver.GetOne("comment", cid, &comment); err == nil {

		db.Driver.Delete(comment)
	} else {
		panic(accountException.AccountCarNotFount())
	}

	ctx.JSON(iris.Map{
		"id": cid,
	})
}
