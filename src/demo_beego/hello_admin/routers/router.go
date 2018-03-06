package routers

import (
	"demo_beego/hello_admin/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
    admin.Run()
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/a", &controllers.MainController{})
	//beego.Router("/a/admin",&rbac.RoleController{}, "*:Index")
}
