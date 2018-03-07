package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init(){
	fmt.Sprint("db_init()  --begin")
	orm.RegisterModel(new(User), new(Product))
	fmt.Sprint("db_init()  --end")
}


func run(){
	init()
}