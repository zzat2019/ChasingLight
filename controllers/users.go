package controllers

import (
	"ChasingLight/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	uuid "github.com/satori/go.uuid"
)

// Operations about Users
type UsersController struct {
	beego.Controller
}

var response = make(map[string]interface{})
var data = make(map[string]interface{})

// @Title Get
// @Description get user by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
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

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (c *UsersController) Logout() {
	c.Data["json"] = "logout successsssssa"
	c.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router /register [post]
func (c *UsersController) Register() {
	var users models.Users
	users.Phone = c.GetString("phone")
	users.Password = c.GetString("password")
	//参数校验
	valid := validation.Validation{}
	valid.Required(users.Phone, "phone")
	valid.Required(users.Phone, "password")
	valid.Phone(users.Phone, "phone")
	if valid.HasErrors() {
		response["code"] = 400
		response["msg"] = "参数错误"
		// 接口成功统一返回
		c.Data["json"] = response
		c.ServeJSON()
	}
	u1, err := uuid.NewV4()
	u2 := uuid.NewV5(u1, "chasinglight").String()
	users.Uuid = u2
	status, err := models.CreateUsers(users.Phone, users.Password, users.Uuid)
	if err != nil {
		response["code"] = 400
		response["msg"] = "创建失败"
		data["status"] = err.Error()
		response["data"] = data
		c.Data["json"] = response
		c.ServeJSON()
	}
	response["code"] = 200
	response["msg"] = "创建成功"
	data["status"] = status
	response["data"] = data
	// 接口成功统一返回
	c.Data["json"] = response
	c.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (c *UsersController) login() {
	var users models.Users
	users.Phone = c.GetString("phone")
	users.Password = c.GetString("password")

}
