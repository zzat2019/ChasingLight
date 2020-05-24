package models

import (
	"github.com/astaxie/beego/orm"
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
	CreateTime int    `orm:"column(create_time)"`
	UpdateTime int    `orm:"column(update_time)"`
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

func CreateUsers(phone string, password string) (id int64, err error) {
	o := orm.NewOrm()
	var users Users
	users.Phone = phone
	users.Password = password
	id, err = o.Insert(&users)
	if err == nil {
		return 8, err
	}
	return 9, err
}
