package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type MainController struct {
	beego.Controller
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 100)
	orm.RegisterModel(new(User))
}

func (ctr *MainController) Get() {
	user := GetUser()
	ctr.Data["json"] = user
	ctr.ServeJSON()
}

func GetUser() User {
	o := orm.NewOrm()
	o.Using("default")

	user := User{
		Id: 1,
	}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		log.Fatal("No result found.")
	} else if err == orm.ErrMissPK {
		log.Fatal("No primary key found.")
	}

	return user
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}