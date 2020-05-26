package util

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var Cache cache.Cache

func Init() {
	collectionName := beego.AppConfig.String("cache.collectionName")
	conn := beego.AppConfig.String("cache.conn")
	dbNum := beego.AppConfig.String("cache.dbNum")
	//password := beego.AppConfig.String("cache.password")
	// 设置配置参数
	config := orm.Params{
		"key":      collectionName,
		"conn":     conn,
		"dbNum":    dbNum,
		"password": "zzat1995",
	}
	configStr, err := json.Marshal(config)
	logs.Debug(string(configStr))
	if err != nil {
		logs.Error("redis配置模型转换失败")
		return
	}
	Cache, err = cache.NewCache("redis", string(configStr))
	if err != nil {
		logs.Error("redis初始化失败")
		return
	}
	logs.Info("******************************************************************************")
	logs.Info("********************************redis启动成功**********************************")
	logs.Info("******************************************************************************")
}
