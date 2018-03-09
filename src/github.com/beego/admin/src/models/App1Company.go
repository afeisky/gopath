package models

import (
	"errors"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
	"fmt"
)

//分组表
type App1Company struct {
	Id     int64
	Name   string  `orm:"size(20)" form:"Name"  valid:"Required"`
	Longname   string  `orm:"size(50)" form:"Longname"  valid:"Required"`
	Email         string    `orm:"null;size(32)" form:"Email" valid:"Email"`
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Status int     `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
}

func (g *App1Company) TableName() string {
	fmt.Println("App1Company....TableName")
	return beego.AppConfig.String("app1_company_table")
}

func init() {
	orm.RegisterModel(new(App1Company))
}

func checkCompany(g *App1Company) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&g)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//get company list
func GetCompanylist(page int64, page_size int64, sort string) (companys []orm.Params, count int64) {
	o := orm.NewOrm()
	company := new(App1Company)
	qs := o.QueryTable(company)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&companys)
	count, _ = qs.Count()
	return companys, count
}

func AddCompany(g *App1Company) (int64, error) {
	if err := checkCompany(g); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	company := new(App1Company)
	company.Name = g.Name
	company.Email = g.Email
	company.Remark = g.Remark
	company.Status = g.Status
	id, err := o.Insert(company)
	return id, err
}

func UpdateCompany(g *App1Company) (int64, error) {
	if err := checkCompany(g); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	company := make(orm.Params)
	if len(g.Name) > 0 {
		company["Name"] = g.Name
	}
	if len(g.Longname) > 0 {
		company["Longname"] = g.Longname
	}
	if len(g.Email) > 0 {
		company["Email"] = g.Email
	}
	if len(g.Remark) > 0 {
		company["Remark"] = g.Remark
	}
	if g.Status != 0 {
		company["Status"] = g.Status
	}
	//if g.Sort != 0 {
	//	company["Sort"] = g.Sort
	//}
	if len(company) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table App1Company
	num, err := o.QueryTable(table).Filter("Id", g.Id).Update(company)
	return num, err
}

func DelCompanyById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&App1Company{Id: Id})
	return status, err
}

func CompanyList() (companys []orm.Params) {
	o := orm.NewOrm()
	company := new(App1Company)
	qs := o.QueryTable(company)
	qs.Values(&companys, "id", "title")
	return companys
}
