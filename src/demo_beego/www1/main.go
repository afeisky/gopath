package main

import (
	_ "demo_beego/www1/routers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
)


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterModel(new(User))

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(123.123.123.123:3306)/test?charset=utf8")
}
func main() {
	fmt.Println("main.go--> START")

	beego.Run()
	fmt.Println("main.go--> END")
}

