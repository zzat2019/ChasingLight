package main

import (
	"ChasingLight/models"
	_ "ChasingLight/routers"
	"ChasingLight/util"
	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	// 想模板中注册函数（首字母大写）
	beego.BConfig.WebConfig.Session.SessionOn = true
	util.Init()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
