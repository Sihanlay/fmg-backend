package views

import (
	_ "context"
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"

	//"encoding/json"
	//"errors"
	//"fmt"
	_ "grpc-demo/exceptions/system"
	_ "grpc-demo/models/db"

	_ "net/http"
	//"strings"
)

func ListAddress(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

}

func CreatAddress(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {

}

func PutAddress(ctx iris.Context, auth authbase.AuthAuthorization, uid int) {
	panic("implement me")
}
