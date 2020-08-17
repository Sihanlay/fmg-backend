package resourceException

import (
	"fmt"
	"grpc-demo/models"
)

func TokenDecodeFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5401,
		Message: "token验证失败",
	}
}

func ModelNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5402,
		Message: "模块不存在",
	}
}

func SaveFileFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5402,
		Message: "保存文件失败",
	}
}

func LenAttachmentsMastSmall5() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5403,
		Message: "附件个数不得超过5",
	}
}

func TestException(i int, ii int) models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5403,
		Message: fmt.Sprintf("%d %d", i, ii),
	}
}
