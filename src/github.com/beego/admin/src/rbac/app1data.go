package rbac

import (
	"fmt"
	"github.com/astaxie/beego"

	"time"
	"encoding/json"
)

type App1DataController struct {
	beego.Controller
}

func (this *App1DataController) Get1() {
	this.Data["json"] = &map[string]interface{}{"success": 0, "data": "bbbb"}
	this.ServeJSON()
}
func (this *App1DataController) Get() {
	//this.Data["data"] =`{"Name":"app1","Time":"2018-03-08","Ver":"v1.0","K1":"key1","K2":"key2","From":"addr","Data":"{\"aaa\":\"bbb\"}"}"}`
	this.Data["data"]=`{"Name":"app1","Time":"2018-03-08","Ver":"v1.0","K1":"key1","K2":"key2","From":"addr","Data":"aa"}`
	this.TplName = "data.tpl"
}
// server send:
type InfoSend struct {
	Success int `json:"success"`
	Time string `json:"time"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
	//K1 string  `json:"data"`    //安全K1: Time加Data共10个字节的MD5码 , 或AES编码
	//K2 string  `json:"data"`    //安全K2: Time加Data共10个字节的MD5码 , 或AES编码
}
// server receive:
type InfoReceiv struct {
	Name string `json:"Name"`
	Time string `json:"Time"`
	Ver string `json:"Ver"`
	//K1 string `json:"K1"` //安全K1: Time加Data共10个字节的MD5码 , 或AES编码
	//K2 string  `json:"K2"`//安全K2: Time加Data共10个字节的MD5码 , 或AES编码
	From string `json:"From"`
	Data interface{} `json:"Data"`
}

func (this *App1DataController) Post() {
	fmt.Println("App1DataController ->POST--->",getNow())
	jsonStr := this.GetString("data")
	//jsonStr := `{"Name":"app1","Time":"2018-03-08","Ver":"v1.0","From":"addr","Data":"aa"}`
	fmt.Println(jsonStr)
	//parse receiver data
	var rev InfoReceiv
	err := json.Unmarshal([]byte(jsonStr), &rev)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = string(outError())
		this.ServeJSON()
		return
	}
	fmt.Println(rev.Name, ":", rev.Time, ":", rev.Data)
	//
	var send *InfoSend = &InfoSend{
		Success: 0,
		Time: "",
		Msg: "success",
		Data:"",
	}
	send.Time=getNow()
	//---
	js, _ := json.Marshal(send)
	fmt.Println(string(js))//{"success":1,"time":""}
	this.Data["json"] = string(js)
	this.ServeJSON()
}
func outError() string {
	var infoServer *InfoSend = &InfoSend{
		Success: 1,
		Time: "",
		Msg: "error",
	}
	infoServer.Time=getNow()
	js, _ := json.Marshal(infoServer)
	fmt.Println(string(js))//{"success":1,"time":""}
	return string(js)
}

func getNow() string { //
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, mon, day, hour, min, sec)
}