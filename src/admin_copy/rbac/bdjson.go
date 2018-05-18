package rbac

import (

	"github.com/astaxie/beego/orm"
	m 	"admin_copy/models"
)

type BdJsonController struct {
	CommonController
}

func (this *BdJsonController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt64("page")
		page_size, _ := this.GetInt64("rows")
		sort := this.GetString("sort")
		order := this.GetString("order")
		if len(order) > 0 {
			if order == "desc" {
				sort = "-" + sort
			}
		} else {
			sort = "Id"
		}
		roles, count := m.GetRolelist(page, page_size, sort)
		if len(roles) < 1 {
			roles = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &roles}
		this.ServeJSON()
		return
	} else {
		this.TplName = this.GetTemplatetype() + "/rbac/role.tpl"
	}

}

func (this *BdJsonController) Getlist() {
	roles, _ := m.GetRolelist(1, 1000, "Id")
	if len(roles) < 1 {
		roles = []orm.Params{}
	}
	this.Data["json"] = &roles
	this.ServeJSON()
	return
}
