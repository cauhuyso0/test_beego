package main

import (
	"fmt"
	"log"
	"test-orm/models"

	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//HeidiSQL
	err := orm.RegisterDataBase("default", "mysql", "root:123456@/beego?charset=utf8")
	if err != nil {
		log.Panicf("dang ky db %v", err)
	}

}

func main() {
	o := orm.NewOrm()

	//user := models.User{Name: "Huy"}

	// id, err := o.Insert(&user)
	// fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// user.Name = "CauHuy"
	// num, err := o.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	// u := models.User{Id: user.Id}
	// err := o.Read(&u)
	// fmt.Printf("ERR: %v\n", err)
	// if num, err := o.Delete(&models.User{Id: 0}); err == nil {
	// 	fmt.Println(num)
	// }

	// fmt.Println(o.Insert(&user))
	// user := models.User{}
	// err := o.Read(&user)
	// if err == orm.ErrNoRows {
	// 	fmt.Println("No result found.")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("No primary key found.")
	// } else {
	// 	fmt.Println(user.Id, user.Name)
	// }

	// đọc dữ liệu bằng câu lệnh hổ trợ
	// var users []models.User
	// o.QueryTable("user").All(&users)

	//đọc dữ liệu trực tiếp bằng câu lệnh SQL
	var users []models.User
	num, err := o.Raw("SELECT id, name FROM user WHERE id=?", 1).QueryRows(&users)
	if err == nil {
		fmt.Println("user nums: ", num, users)
	}

}
