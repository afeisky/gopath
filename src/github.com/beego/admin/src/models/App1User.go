package models

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	. "github.com/beego/admin/src/lib"
	"fmt"
)

//用户表
type ProductUser struct {
	Id            int64
	Username      string    `orm:"size(20)" form:"Code" valid:"Required;MaxSize(20);MinSize(6)"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Nickname      string    `orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Ckey      string    `orm:"size(16)" form:"Ckey" valid:"Required;MaxSize(20);MinSize(6)"`
	//ProductCompany          []*ProductCompany   `orm:"rel(m2m)"`
	CompanyId    int64   `form:"CompanyId"  valid:"Required"`
}

func (u *ProductUser) TableName() string {
	fmt.Println("ProductUser....TableName")
	return beego.AppConfig.String("app1_user_table")
}

func (u *ProductUser) Valid(v *validation.Validation) {
	fmt.Println("ProductUser....Valid()")
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkProductUser(u *ProductUser) (err error) {
	fmt.Println("ProductUser....checkProductUser()")
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	orm.RegisterModel(new(ProductUser))
}

/************************************************************/

//get user list
func GetProductUserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	fmt.Println("ProductUser....GetProductUserlist()",page,",",page_size,",",sort,",count=",count)
	o := orm.NewOrm()
	user := new(ProductUser)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}

//添加用户
func AddProductUser(u *ProductUser) (int64, error) {
	fmt.Println("ProductUser....AddProductUser()")
	if err := checkProductUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(ProductUser)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status
	user.CompanyId = u.CompanyId
	id, err := o.Insert(user)
	return id, err
}

//更新用户
func UpdateProductUser(u *ProductUser) (int64, error) {
	fmt.Println("ProductUser....UpdateProductUser()")
	if err := checkProductUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Nickname) > 0 {
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Remark) > 0 {
		user["Remark"] = u.Remark
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if u.CompanyId != 0 {
		user["CompanyId"] = u.CompanyId
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table ProductUser
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func DelProductUserById(Id int64) (int64, error) {
	fmt.Println("ProductUser....DelProductUserById()")
	o := orm.NewOrm()
	status, err := o.Delete(&ProductUser{Id: Id})
	return status, err
}

func GetProductUserByUsername(username string) (user ProductUser) {
	fmt.Println("ProductUser....GetProductUserByUsername()")
	user = ProductUser{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

func GetProductUserById(id int64) (user ProductUser) {
	fmt.Println("ProductUser....GetProductUserById()")
	user = ProductUser{Id: id}
	o := orm.NewOrm()
	o.Read(&user, "Id")
	return user
}

//-----------------------------------------------
func GetCompanyList() (companys []orm.Params, count int64) {
	o := orm.NewOrm()
	company := new(ProductCompany)
	count, _ = o.QueryTable(company).Values(&companys)
	return companys, count
}
func GetCompanyIdByUserId(userid int64) (companyid int64) {
	o := orm.NewOrm()
	var table ProductUser
	o.QueryTable(table).Filter("Id", userid).One(&table)
	companyid=table.CompanyId
	fmt.Println("GetCompanyIdById....companyid=",companyid)
	return companyid
}
//更新用户
func UserUpdateCompany(userid int64,companyid int64) (int64, error) {
	fmt.Println("ProductUser....UpdateProductUser()")
	o := orm.NewOrm()
	var table ProductUser
	user := make(orm.Params)
	user["CompanyId"] = companyid
	num, err := o.QueryTable(table).Filter("Id", userid).Update(user)
	return num, err
}
/*
func (this *TestController) Show() {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("select arc.id,arc.title,arc.typeid,art.typename from go_archives as arc left join go_arctype as art on art.id=arc.typeid where arc.typeid=?", 2).Values(&maps)
	fmt.Println(maps)
}

func (this *TestController) Show3() {
	o := orm.NewOrm()
	var arc []*models.Archives
	o.QueryTable("go_archives").Filter("Id__gt", 1).RelatedSel().All(&arc)//使用RelatedSel将关联的arctype也查出来，也就是left join arctype as T1 on T1.id=go_archives.arctype_id
	arc3 := arc[0]
	fmt.Println(arc3.Arctype.Typename)
}
*/