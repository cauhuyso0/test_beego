package models

import (
	"github.com/astaxie/beego/client/orm"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}
