package db

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/shoogoome/mutils/hash"
	paramsUtils "github.com/shoogoome/mutils/params"
	reflectUtils "github.com/shoogoome/mutils/reflect"
	"grpc-demo/constants"
	"grpc-demo/core/cache"
	systemException "grpc-demo/exceptions/system"
	"grpc-demo/utils"

	"reflect"
	"regexp"
	"strings"
	"time"
)

var Driver driver

func InitDB() {

Conn:
	o, err := gorm.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
			utils.GlobalConfig.Mysql.Username,
			utils.GlobalConfig.Mysql.Password,
			utils.GlobalConfig.Mysql.Host,
			utils.GlobalConfig.Mysql.Port,
			utils.GlobalConfig.Mysql.DB,
		))
	Driver.DB = o
	if err != nil {
		fmt.Println("[!] 数据库链接异常，尝试重新链接...", err)
		time.Sleep(time.Second * 5)
		goto Conn
	}
	Driver.Callback().Create().Before("gorm:create").Register("update_time_in_create", updateTimeForCreateCallback)
	Driver.Callback().Update().Before("gorm:update").Register("update_time_in_update", updateTimeForUpdateCallback)
	Driver.Callback().Update().After("gorm:update").Register("delete_cache_after_update", deleteCacheSignal)
	Driver.Callback().Delete().After("gorm:delete").Register("delete_cache_after_update", deleteCacheSignal)
	Driver.DB.DB().SetMaxIdleConns(30)
	Driver.DB.DB().SetMaxOpenConns(500)
	Driver.SingularTable(true)

	Driver.AutoMigrate(
		&Account{}, &Address{}, &AccountCar{}, &Delivery{},
		&Comment{}, &Like{},&News{},&NewsAndTag{},&NewsTag{},
	)
}

// 更新创建时间
func updateTimeForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if scope.HasColumn("CreateTime") {
			scope.SetColumn("CreateTime", time.Now().Unix())
		}
		updateTimeForUpdateCallback(scope)
	}
}

// 更新更新时间
func updateTimeForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		// 更新时间
		if scope.HasColumn("UpdateTime") {
			scope.SetColumn("UpdateTime", time.Now().Unix())
		}
	}
}

// 删除缓存
func deleteCacheSignal(scope *gorm.Scope) {
	if !scope.HasError() {
		id, ok := scope.FieldByName("id")
		if ok && reflectUtils.IsExist(id.Field) {
			key := paramsUtils.CacheBuildKey(constants.DbModel, scope.TableName(), id.Field.Int())
			cache.Redis.Do(constants.DbNumberModel, "del", key)
		}
	}
}

type driver struct {
	*gorm.DB
}

func (d *driver) Exec(sql string, values ...interface{}) *gorm.DB {

	var ids []struct {
		Id int
	}

	// 清除缓存
	updateCompile := regexp.MustCompile("^update.+`(.*)`.+where(.*)$")
	deleteCompile := regexp.MustCompile("^delete.+from.+`(.*)`.+where(.*)$")
	sqlSplit := strings.Split(sql, " ")

	if len(sqlSplit) > 0 {
		var r [][]byte
		switch sqlSplit[0] {
		case "update":
			r = updateCompile.FindSubmatch([]byte(sql))

		case "delete":
			r = deleteCompile.FindSubmatch([]byte(sql))
		}
		if len(r) == 3 {
			table := string(r[1])
			search := string(r[2])
			if len(values) > 0 {
				quire := values[len(values)-strings.Count(string(r[2]), "?"):]
				d.DB.Raw(fmt.Sprintf("select id from `%s` where %s", table, search), quire...).Scan(&ids)
			} else {
				d.DB.Raw(fmt.Sprintf("select id from `%s` where %s", table, search)).Scan(&ids)
			}
			for _, i := range ids {
				// 清除model缓存
				key := paramsUtils.CacheBuildKey(constants.DbModel, table, i.Id)
				cache.Redis.Do(constants.DbNumberModel, "del", key)
			}
		}
	}

	db := d.DB.Exec(sql, values...)
	defer func() {
		if err := recover(); err != nil {
			panic(systemException.SystemException())
		}
	}()
	return db
}

// 走缓存获取一条记录
func (d *driver) GetOne(table string, id int, target interface{}, db ...*gorm.DB) error {

	key := paramsUtils.CacheBuildKey(constants.DbModel, table, id)
	object, err := redis.Bytes(cache.Redis.Do(constants.DbNumberModel, "get", key))

	if err == nil && object != nil {
		err := json.Unmarshal(object, &target)
		return err
	}

	myValue := reflect.ValueOf(target)
	myType := reflect.TypeOf(target)

	x := reflect.New(myType)
	x.Elem().Set(myValue)

	if len(db) > 0 {
		err = db[0].First(x.Interface(), id).Error
	} else {
		err = d.First(x.Interface(), id).Error
	}
	if err != nil {
		return err
	} else {
		target = x.Interface()
		if resByte, err := json.Marshal(&target); err == nil {
			v := hash.RandInt64(240, 240*5)
			_, _ = cache.Redis.Do(constants.DbNumberModel, "set", key, resByte, int(v)*60*60)
		} else {
			return err
		}
	}
	return nil
}

// 走缓存获取多条记录
func (d *driver) GetMany(table string, ids []interface{}, target interface{}, db ...*gorm.DB) []interface{} {
	myValue := reflect.ValueOf(target)
	myType := reflect.TypeOf(target)

	data := make([]interface{}, 0, len(ids))

	for _, id := range ids {
		x := reflect.New(myType)
		x.Elem().Set(myValue)
		if err := d.GetOne(table, int(id.(float64)), x.Interface(), db...); err == nil {
			data = append(data, x.Elem().Interface())
		}
	}
	return data
}
