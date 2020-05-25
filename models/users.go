package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct {
	Id         int64  `orm:"column(id)"`
	Uuid       string `orm:"column(uuid)"`
	WxOpenid   string `orm:"column(wx_openid)"`
	AliUid     string `orm:"column(ali_uid)"`
	Icon       string `orm:"column(icon)"`
	NickName   string `orm:"column(nick_name)"`
	Phone      string `orm:"column(phone)"`
	Gender     int    `orm:"column(gender)"`
	Password   string `orm:"column(password)"`
	SystemStar int    `orm:"column(system_star)"`
	OwnStar    int    `orm:"column(own_star)"`
	CreateTime int64  `orm:"column(create_time)"`
	UpdateTime int64  `orm:"column(update_time)"`
	Status     int    `orm:"column(status)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func (t *Users) TableName() string {
	return "user"
}

func GetUsers(id int64) (s Users, err error) {
	o := orm.NewOrm()
	o.Using("user")
	s = Users{Id: id}
	err = o.Read(&s)

	if err != nil {
		return s, err
	}
	return
}

func CreateUsers(phone string, password string, uuid string) (status bool, err error) {
	o := orm.NewOrm()
	var users Users
	var maps []orm.Params
	users.Phone = phone
	users.Password = password
	users.Uuid = uuid
	res, err := o.Raw("select uuid from user where phone = ?", phone).Values(&maps)
	if err != nil || res > 0 {
		return false, errors.New("该手机号已被注册")
	}
	users.CreateTime = time.Now().Unix()
	users.UpdateTime = time.Now().Unix()
	users.Status = 1
	id, err := o.Insert(&users)
	_ = id
	if err != nil {
		return false, errors.New("插入失败")
	}
	return true, nil
}
