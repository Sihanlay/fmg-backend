package resourceLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"grpc-demo/constants"
	"grpc-demo/exceptions/resource"
	"grpc-demo/models"
	"grpc-demo/utils/hash"
)

type ResourcesLocalStorage struct {
	Model string
	root  string
}

type tokenStruct struct {
	ExpireAt  int64  `json:"expire_at"`
	AccountId int    `json:"account_id"`
	Path      string `json:"path"`
	Public    bool   `json:"public"`
	FileName  string `json:"file_name"`
}

type TokenOption struct {
	FileName   string
	Encryption bool
}

func NewReousrcesLocalStorage(model string) ResourcesLocalStorage {
	if _, ok := constants.StorageMapping[model]; !ok {
		panic(resourceException.ModelNotExists())
	}
	return ResourcesLocalStorage{
		Model: model,
		root:  constants.StorageMapping[model],
	}
}

func (r *ResourcesLocalStorage) SaveFile(npath string, file []byte, cache bool) string {
	dirName, fileName := path.Split(npath)
	_ = os.MkdirAll(path.Join(r.root, dirName), 0777)

	var err error
	fileWrite := &os.File{}
	if cache {
		npath = path.Join(dirName, fmt.Sprintf("%d-%s", time.Now().Unix(), fileName))
	}
	fileWrite, err = os.OpenFile(path.Join(r.root, npath), os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		panic(resourceException.SaveFileFail())
	}
	defer fileWrite.Close()
	_, err = io.Copy(fileWrite, bytes.NewReader(file))
	if err != nil {
		panic(resourceException.SaveFileFail())
	}
	return fmt.Sprintf("storage://%s@%s", r.Model, npath)
}

// 批量上传附件
func (r *ResourcesLocalStorage) UploadAttachments(ctx iris.Context, resources string, _id int) (string, iris.Map) {
	var dict []models.FileFormat
	json.Unmarshal([]byte(resources), &dict)

	count := len(dict)
	// 卡附件个数
	if count >= 5 {
		panic(resourceException.LenAttachmentsMastSmall5())
	}

	fileId := 0
	fileNameList := make([]string, len(dict))
	for i, v := range dict {
		fileNameList[i] = v.Name
		if v.Id > fileId {
			fileId = v.Id
		}
	}
	// 解析读取模式
	ctx.Request().ParseMultipartForm(ctx.Application().ConfigurationReadOnly().GetPostMaxMemory())
	filesList := ctx.Request().MultipartForm.File

	fileId += 1
	status := iris.Map{}
	npath := fmt.Sprintf("%s/%d", r.root, _id)
	_ = os.MkdirAll(npath, 0777)
	for _, files := range filesList {
		for _, file := range files {
			// 超过5个附件
			if count >= 5 {
				status[file.Filename] = -1
				continue
			}
			// 去除超重文件
			if file.Size > 1024*1024*10 {
				status[file.Filename] = 0
				continue
			}
			// 文件重名处理
			index := 1
			fileNameSplit := strings.Split(file.Filename, ".")
			fileNamePrefix := strings.Join(fileNameSplit[:len(fileNameSplit)-1], "")
			fileNameSuffix := fileNameSplit[len(fileNameSplit)-1]
			newFileName := file.Filename
			for stringIsExists(newFileName, fileNameList) {
				newFileName = fmt.Sprintf("%s(%d).%s", fileNamePrefix, index, fileNameSuffix)
				index += 1
			}

			// 元数据记录
			token := hash.GetRandomString(16)
			dict = append(dict, models.FileFormat{
				Id:         fileId,
				Name:       newFileName,
				Storage:    fmt.Sprintf("storage://%s@%s", r.Model, fmt.Sprintf("%d/%s", _id, token)),
				Size:       file.Size,
				CreateTime: time.Now().Unix(),
			})
			fileId += 1
			file.Filename = token
			uploadTo(file, npath)
			status[newFileName] = 1
			count += 1
		}
	}
	resourcesByte, _ := json.Marshal(&dict)
	return string(resourcesByte), status
}

// 删除附件
func (r *ResourcesLocalStorage) DeleteAttachments(resources string, ids []interface{}) string {

	var dict []models.FileFormat

	json.Unmarshal([]byte(resources), &dict)

	newResources := make([]models.FileFormat, 0, len(dict))
	deleteResources := make([]models.FileFormat, 0, len(dict))

	// 区分删除和保留资源
	for _, i := range dict {
		if intIsExists(i.Id, ids) {
			deleteResources = append(deleteResources, i)
		} else {
			newResources = append(newResources, i)
		}
	}
	// 删除物理资源
	for _, i := range deleteResources {
		re := regexp.MustCompile(`^storage://(.*)@(.*)$`)
		r := re.FindSubmatch([]byte(i.Storage))
		os.RemoveAll(path.Join(constants.StorageMapping[string(r[1])], string(r[2])))
	}
	resource, _ := json.Marshal(&newResources)
	return string(resource)
}

// 生成下载token
func GenerateToken(path string, aid int, expire int64, option ...TokenOption) string {
	if len(path) == 0 {
		return ""
	}
	public := false
	if aid == -1 {
		public = true
	}
	payload := tokenStruct{
		ExpireAt:  -1,
		AccountId: aid,
		Path:      path,
		Public:    public,
	}
	if len(option) > 0 {
		payload.FileName = option[0].FileName
		if option[0].Encryption {
			payload.ExpireAt = expire + time.Now().Unix()
		}
		return hash.GenerateToken(payload, option[0].Encryption)
	}
	return hash.GenerateToken(payload)
}

// 解析token
func DecodeToken(token string, aid int) (string, string, string) {
	// 拦截解码方法报错，传递自身报错
	defer func() {
		if err := recover(); err != nil {
			panic(resourceException.TokenDecodeFail())
		}
	}()
	// 反序列化
	var resourceToken tokenStruct
	hash.DecodeToken(token, &resourceToken)
	// 鉴权
	if !resourceToken.Public && resourceToken.AccountId != aid {
		panic(resourceException.TestException(aid, resourceToken.AccountId))
	}
	// 检查有效性
	if resourceToken.ExpireAt > 0 && resourceToken.ExpireAt < time.Now().Unix() {
		panic(resourceException.TokenDecodeFail())
	}
	// 路径匹配
	re := regexp.MustCompile(`^storage://(.*)@(.*)$`)
	r := re.FindSubmatch([]byte(resourceToken.Path))
	if len(r) != 3 {
		panic(resourceException.TokenDecodeFail())
	}
	fpath, fileName := path.Split(
		path.Join(constants.StorageMapping[string(r[1])],
			string(r[2])))
	return fpath, fileName, resourceToken.FileName
}

// 保存文件
func uploadTo(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}

func stringIsExists(s string, sl []string) bool {
	for _, i := range sl {
		if s == i {
			return true
		}
	}
	return false
}

func intIsExists(s int, sl []interface{}) bool {
	for _, i := range sl {
		if s == int(i.(float64)) {
			return true
		}
	}
	return false
}
