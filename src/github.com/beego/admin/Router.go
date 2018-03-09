package admin

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin/src/rbac"
)

func router() {
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

	//BEGIN-wuchaofei add
	beego.Router("/app1/company/AddCompany", &rbac.App1CompanyController{}, "*:AddCompany")
	beego.Router("/app1/company/UpdateCompany", &rbac.App1CompanyController{}, "*:UpdateCompany")
	beego.Router("/app1/company/DelCompany", &rbac.App1CompanyController{}, "*:DelCompany")
	beego.Router("/app1/company/index", &rbac.App1CompanyController{}, "*:Index")

	beego.Router("/app1/user/AddUser", &rbac.App1UserController{}, "*:AddUser")
	beego.Router("/app1/user/UpdateUser", &rbac.App1UserController{}, "*:UpdateUser")
	beego.Router("/app1/user/DelUser", &rbac.App1UserController{}, "*:DelUser")
	beego.Router("/app1/user/index", &rbac.App1UserController{}, "*:Index")
	beego.Router("/app1/user/UserSelectCompanyList", &rbac.App1UserController{}, "*:UserSelectCompanyList")
	beego.Router("/app1/user/UserUpdateCompany", &rbac.App1UserController{}, "*:UserUpdateCompany")
	beego.Router("/app1/data", &rbac.App1DataController{},"get:Get")
	beego.Router("/app1/data", &rbac.App1DataController{},"post:Post")

	//END---wuchaofei add
}
