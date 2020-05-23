package controllers

import (
	"ChasingLight/models"
	"errors"
	"github.com/astaxie/beego"
)

// Operations about Users
type UsersController struct {
	beego.Controller
}

var response = make(map[string]interface{})

func (c *UsersController) Get() {
	var users models.Users
	err := func() error {
		id, err := c.GetInt64("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		users, err = models.GetUsers(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err != nil {
		c.Ctx.WriteString("wrong params")
	}

	response["code"] = 200
	response["msg"] = "查询成功"
	response["data"] = users
	// 接口成功统一返回
	c.Data["json"] = response
	c.ServeJSON()
}
