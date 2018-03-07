package rbac

import (
	m "github.com/beego/admin/src/models"

	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

type ProductUserController struct {
	CommonController
}

func (this *ProductUserController) Index() {
	fmt.Println("ProductUserController....Index()")
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
	users, count := m.GetProductUserlist(page, page_size, sort)
	if this.IsAjax() {
		fmt.Println("ProductUserController....1,",users)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
		return
	} else {
		fmt.Println("ProductUserController....2,",users)
		tree := this.GetTree()
		this.Data["tree"] = &tree
		this.Data["users"] = &users
		if this.GetTemplatetype() != "easyui" {
			this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
		}
		this.TplName = this.GetTemplatetype() + "/app1/user.tpl"
	}

}

func (this *ProductUserController) AddUser() {
	fmt.Println("ProductUserController....AddUser()")
	u := m.ProductUser{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddProductUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProductUserController) UpdateUser() {
	fmt.Println("ProductUserController....UpdateUser()")
	u := m.ProductUser{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateProductUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProductUserController) DelUser() {
	fmt.Println("ProductUserController....DelUser()")
	Id, _ := this.GetInt64("Id")
	status, err := m.DelProductUserById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}

func (this *ProductUserController) UserSelectCompanyList() {
	fmt.Println("ProductUserController UserSelectCompanyList()")
	userid, _ := this.GetInt64("Id")
	if this.IsAjax() {
		companys, count := m.GetCompanylist(1, 1000, "Id")
		companyid := m.GetCompanyIdByUserId(userid)
		for i := 0; i < len(companys); i++ {
			if companys[i]["Id"] == companyid {
				companys[i]["checked"] = 1
			}
		}
		if len(companys) < 1 {
			companys = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &companys}
		this.ServeJSON()
		return
	} else {
		this.Data["userid"] = userid
		this.TplName = this.GetTemplatetype() + "/app1/userselectcompanylist.tpl"
	}
}

func (this *ProductUserController) UserUpdateCompany() {
	fmt.Println("ProductUserController --UserUpdateCompany()")
	userid, _ := this.GetInt64("Id")
	companyids := this.GetString("ids")
	userids := strings.Split(companyids, ",")
	if len(companyids) > 0 {
		for _, v := range userids {
			id, _ := strconv.Atoi(v)
			_, err := m.UserUpdateCompany(userid, int64(id))
			if err != nil {
				this.Rsp(false, err.Error())
			}
		}
	}
	this.Rsp(true, "success")
}