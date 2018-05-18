package models

import (
	"errors"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

//角色表
type BdJson struct {
	Id int64
	Create_time time.Time `orm:"null;type(datetime)" form:"-"`  // `create_time` datetime NOT NULL,
	Filename   string  `orm:"size(50)"` //`filename` varchar(50) NOT NULL DEFAULT '',
	//type   string  `orm:"size(10)" form:"type"  valid:"Required"` //`type` varchar(10) DEFAULT NULL,
	Imei   string  `orm:"size(10)"` //`imei` varchar(16) DEFAULT '',
	Time time.Time `orm:"null;type(datetime)" form:"-"`  //`time` datetime DEFAULT NULL,
	Comment1  string  `orm:"size(50)"` //`comment1` varchar(50) DEFAULT '',
	Data_ids int64   //`data_ids` int(11) DEFAULT NULL,
	Json   orm.JsonbField   //`json` json DEFAULT NULL,
	Filesave  orm.JsonbField    //`filesave` json DEFAULT NULL,
	Done int64  //`done` int(11) DEFAULT '0',
	Done_time time.Time `orm:"null;type(datetime)" form:"-"` //`done_time` datetime DEFAULT NULL,
}


func (r *BdJson) TableName() string {
	return beego.AppConfig.String("rbac_BdJson_table")
}

func init() {
	orm.RegisterModel(new(BdJson))
}

func checkBdJson(g *BdJson) (err error) {
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

//get BdJson list
func GetBdJsonlist(page int64, page_size int64, sort string) (bdjsons []orm.Params, count int64) {
	o := orm.NewOrm()
	bdJson := new(BdJson)
	qs := o.QueryTable(bdJson)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&bdjsons)
	count, _ = qs.Count()
	return bdjsons, count
}