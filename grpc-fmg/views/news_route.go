package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/study_info"
)

func RegisterNewsRouters(app *iris.Application) {
	NewsRouter := app.Party("news/info")
	NewsRouter.Post("/create", hero.Handler(study_info.CreatStudyInfo))
	NewsRouter.Delete("/del/{nid:int}", hero.Handler(study_info.DeleteStudyInfo))
	NewsRouter.Get("/list", hero.Handler(study_info.ListStudyInfos))
	NewsRouter.Put("/put/{nid:int}", hero.Handler(study_info.PutStudyInfo))
	NewsRouter.Post("/_mget", hero.Handler(study_info.MgetStudyInfo))
}
