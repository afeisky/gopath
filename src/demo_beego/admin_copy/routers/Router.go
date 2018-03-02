package router

import (
	"demo_beego/admin_copy/controllers"
	"github.com/astaxie/beego"
	"fmt"
	"mime"
	"os"
	. "demo_beego/admin_copy/lib"
	"demo_beego/admin_copy/models"
	"demo_beego/admin_copy/rbac"
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
	//router()

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
func router() {
	Run()
	beego.Router("/", &rbac.MainController{}, "*:Index")
	beego.Router("/public/index", &rbac.MainController{}, "*:Index")
	beego.Router("/public/login", &rbac.MainController{}, "*:Login")
	beego.Router("/public/logout", &rbac.MainController{}, "*:Logout")
	beego.Router("/public/changepwd", &rbac.MainController{}, "*:Changepwd")

	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/user/index", &rbac.UserController{}, "*:Index")

	beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	beego.Router("/rbac/node/index", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group/index", &rbac.GroupController{}, "*:Index")

	beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	beego.Router("/rbac/role/Getlist", &rbac.RoleController{}, "*:Getlist")
	beego.Router("/rbac/role/index", &rbac.RoleController{}, "*:Index")
	//-----
	beego.Router("/", &controllers.MainController{})
	beego.AddFuncMap("stringsToJson", StringsToJson)
}
