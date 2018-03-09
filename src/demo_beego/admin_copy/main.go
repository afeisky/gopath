package main

import (
	_ "demo_beego/admin_copy/routers"
	"github.com/astaxie/beego"
	. "demo_beego/admin_copy/lib"
	"fmt"
	"mime"
	"os"
	"demo_beego/admin_copy/models"
)

const VERSION = "0.1.1"

func Run() {
	//初始化
	initialize()

	fmt.Println("Starting....")

	fmt.Println("Start ok")
}
func initialize() {
	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()
	models.Connect()
}
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}
func main() {
	beego.AddFuncMap("stringsToJson", StringsToJson)
	Run()
	beego.Run()
}

