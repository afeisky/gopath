package rbac

import (
	m "github.com/beego/admin/src/models"

)



type ProductCompanyController struct {
	CommonController
}

func (this *ProductCompanyController) Index() {
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
		nodes, count := m.GetCompanylist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
		this.ServeJSON()
		return
	} else {
		this.TplName = this.GetTemplatetype() + "/app1/company.tpl"
	}

}
func (this *ProductCompanyController) AddCompany() {
	g := m.ProductCompany{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddCompany(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProductCompanyController) UpdateCompany() {
	g := m.ProductCompany{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateCompany(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProductCompanyController) DelCompany() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelCompanyById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}

