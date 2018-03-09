package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"fmt"
)

//用户表
type App1Data struct {
	Id            int64
	CId    int64   `form:"CId"  valid:"Required"`
	Time          time.Time `orm:"null;type(datetime)" form:"-"`
	X    	int64   `form:"X"  valid:"Required"`
	Y    	int64   `form:"Y"  valid:"Required"`
}

func (u *App1Data) TableName() string {
	fmt.Println("App1Data....TableName")
	return beego.AppConfig.String("app1_data_table")
}

func (u *App1Data) Valid(v *validation.Validation) {
	//if u.Password != u.Repassword {
	//	v.SetError("Repassword", "两次输入的密码不一样")
	//}
}

func init() {
	orm.RegisterModel(new(App1Data))
}

/************************************************************/
//添加Data
func AddApp1Data(u *App1Data) (int64, error) {
	//if err := checkApp1Client(u); err != nil {
	//	return 0, err
	//}
	o := orm.NewOrm()
	data := new(App1Data)
	data.Time = u.Time
	data.X = u.X
	data.Y = u.Y
	id, err := o.Insert(data)
	return id, err
}

