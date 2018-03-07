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
type App1Client struct {
	Id            int64
	Clientname      string    `orm:"size(20)" form:"Name" valid:"Required;MaxSize(20);MinSize(6)"`
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
	UserId    int64   `form:"UserId"  valid:"Required"`
}

func (u *App1Client) TableName() string {
	fmt.Println("App1Client....TableName")
	return beego.AppConfig.String("app1_client_table")
}

func (u *App1Client) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkApp1Client(u *App1Client) (err error) {
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
	orm.RegisterModel(new(App1Client))
}

/************************************************************/

//get client list
func GetApp1Clientlist(page int64, page_size int64, sort string) (clients []orm.Params, count int64) {
	o := orm.NewOrm()
	client := new(App1Client)
	qs := o.QueryTable(client)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&clients)
	count, _ = qs.Count()
	return clients, count
}

//添加用户
func AddApp1Client(u *App1Client) (int64, error) {
	if err := checkApp1Client(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	client := new(App1Client)
	client.Clientname = u.Clientname
	client.Password = Strtomd5(u.Password)
	client.Nickname = u.Nickname
	client.Email = u.Email
	client.Remark = u.Remark
	client.Status = u.Status

	id, err := o.Insert(client)
	return id, err
}

//更新用户
func UpdateApp1Client(u *App1Client) (int64, error) {
	if err := checkApp1Client(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	client := make(orm.Params)
	if len(u.Clientname) > 0 {
		client["clientname"] = u.Clientname
	}
	if len(u.Nickname) > 0 {
		client["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		client["Email"] = u.Email
	}
	if len(u.Remark) > 0 {
		client["Remark"] = u.Remark
	}
	if len(u.Password) > 0 {
		client["Password"] = Strtomd5(u.Password)
	}
	if u.Status != 0 {
		client["Status"] = u.Status
	}
	if len(client) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table App1Client
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(client)
	return num, err
}

func DelApp1ClientById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&App1Client{Id: Id})
	return status, err
}

func GetApp1ClientByClientname(clientname string) (client App1Client) {
	client = App1Client{Clientname: clientname}
	o := orm.NewOrm()
	o.Read(&client, "clientname")
	return client
}

func GetApp1ClientById(id int64) (client App1Client) {
	client = App1Client{Id: id}
	o := orm.NewOrm()
	o.Read(&client, "Id")
	return client
}

