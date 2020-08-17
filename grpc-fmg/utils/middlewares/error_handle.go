package middlewares

import (
	"grpc-demo/models"
	logUtils "grpc-demo/utils/log"
	"fmt"
	"github.com/kataras/iris"
	"runtime"
)

// 异常控制器
func AbnormalHandle(ctx iris.Context) {
	defer func() {
		re := recover()
		if re == nil {
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		// 打印堆栈信息
		log := fmt.Sprintf("%v\n%s", re, stack())
		if debug, err := ctx.URLParamInt("debug"); err == nil && debug == 1 {
			ctx.Text(log)
			return
		}
		// 输出api格式反馈
		switch result := re.(type) {
		case models.RestfulAPIResult:
			ctx.JSON(result)
		default:
			ctx.JSON(models.RestfulAPIResult{
				Status: false,
				ErrCode: 500,
				Message: fmt.Sprintf("系统错误: %v", result),
			})
			logUtils.Println(log)
		}
	}()
	ctx.Next()
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}

