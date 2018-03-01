package routers

import (
	"demo_beego/www1/controllers"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"time"
	"os"
)

type HelloController struct {
	beego.Controller
}
func (this *HelloController) Get() {
	this.Ctx.WriteString("hello world")
}
func (this *HelloController) Post() {
	_, h, err := this.GetFile("file")
	if err != nil {
		fmt.Println("getfile err ", err)
	}
	fmt.Println("filename:", h.Filename)
	this.SaveToFile("file", "e://temp/upload/"+h.Filename)
	this.Ctx.WriteString("hello 1111")
}
type JsonController struct {
	beego.Controller
}

func (this *JsonController) Get() {
	this.Data["data"] = `{"success": 0, "msg": "111","data":"afsdfsf"}`
	this.TplName =  "a.tpl"


}

func outError() string {
	var infoServer *InfoServer = &InfoServer{
		Success: 1,
		Time: "",
		Msg: "error",
	}
	infoServer.Time=getNow()
	js, _ := json.Marshal(infoServer)
	fmt.Println(string(js))//{"success":1,"time":""}
	return string(js)
}

func (this *JsonController) Post() {
	var infoServer *InfoServer = &InfoServer{
		Success: 1,
		Time: "",
		Msg: "",
	}
	infoServer.Time=getNow()
	//---

	jsonStr := this.GetString("data")
	fmt.Println("POST--->",getNow(),jsonStr)
	jsonData := []byte(jsonStr)
	//jsonStr := []byte(`{"rd":"test","cmd":"001001","data":{"a":"aaa"},"key":"001001","time":"2018-03-01 11:38:54"}`)
	var infoRequest InfoRequest
	var f interface{}
	err := json.Unmarshal(jsonData, &f)
	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString(outError())
		return
	}
	jsonObjectParse(f)
	var ctime string
	var rd string
	var data interface{}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "rd":
			rd=v.(string)
			infoRequest.App=v.(string)
		case "time":
			ctime=v.(string)
			infoRequest.Time=v.(string)
		case "data":
			data=v
			infoRequest.Data=v
		default:
			fmt.Println(k, "[type?]",v)
		}
	}
	fmt.Println("rd",rd)
	fmt.Println("ctime",ctime)
	fmt.Println("data",data)

	if !checkRequest(infoRequest){
		fmt.Println("Error:",infoRequest.App)
	}
	infoServer.Time=getNow()
	infoServer.Success=0
	js, _ := json.Marshal(infoServer)
	fmt.Println(string(js))//{"success":1,"time":""}
	this.Ctx.WriteString(string(js))
	//this.Data["data"] =jsonStr
	//this.TplName =  "a.tpl"
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}


func jsonArrayParse(vv []interface{}){
	for i, u := range vv {
		//var f1 interface{}
		switch vv1 := u.(type) {
		case string:
			fmt.Println(i, "[string_] :", u)
		case float64:
			fmt.Println(i, "[float64_]:", u)
		case bool:
			fmt.Println(i, "[bool_]:", u)
		case nil:
			fmt.Println(i, "[nil_]:", u)
		case []interface{}:
			fmt.Println(i, "[array_] :", u)
			jsonArrayParse(vv1)
		case interface{}:
			fmt.Println(i, "[interface_]:",u)
			m1 := u.(map[string]interface{})
			jsonObjectParse(m1)
		default:
			fmt.Println("  ", i, "[type?_]", u, ", ",vv1)
		}
	}
}

func jsonObjectParse(f interface{}){
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "[string] :", vv)
		case float64:
			fmt.Println(k, "[float64]:", vv)
		case bool:
			fmt.Println(k, "[bool]:", vv)
		case nil:
			fmt.Println(k, "[nil]:", vv)
		case []interface{}:
			fmt.Println(k, "[array]:")
			jsonArrayParse(vv)
		case interface{}:
			fmt.Println(k, "[interface]:",vv)
			m1 := v.(map[string]interface{})
			jsonObjectParse(m1)
		default:
			fmt.Println(k, "[type?]",vv)
		}
	}
}

func nowTime111() string {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	return time.Now().Format("2006-01-02 15:04:05")
}
func getNow() string { //��ȡ��ǰʱ��
	//    now := time.Now()
	//    year, mon, day := now.UTC().Date()
	//    hour, min, sec := now.UTC().Clock()
	//    zone, _ := now.UTC().Zone()
	//    fmt.Printf("UTC time is %d-%d-%d %02d:%02d:%02d %s\n",
	//        year, mon, day, hour, min, sec, zone)
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	//zone, _ = now.Zone()
	//    fmt.Printf("local time is %d-%d-%d %02d:%02d:%02d %s\n",
	//        year, mon, day, hour, min, sec, zone)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, mon, day, hour, min, sec)
}


type InfoRequest struct {
	App string
	Time string
	K1 string  //��ȫK1 ΪTime��Dataǰ10λ�ϳɵĴ����ɵ�MD5�� , ��AES��
	K2 string  //��ȫK2 ΪTime��Data���кϳɵĴ����ɵ�MD5��
	From string
	Data interface{}
}
func checkRequest(info InfoRequest) bool{
	if (info.App=="test") {
		return true
	}
	return false
}
type InfoServer struct {
	Success int `json:"success"`
	Time string `json:"time"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
	K1 string  `json:"data"`    //��ȫK1 ΪTime��Dataǰ10λ�ϳɵĴ����ɵ�MD5�� , ��AES��
	K2 string  `json:"data"`    //��ȫK2 ΪTime��Data���кϳɵĴ����ɵ�MD5��
}

func init() {
	fmt.Println("Hello Beego")
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &HelloController{})
	beego.Router("/a", &JsonController{})
}

//--byte ת string
// strByte := []byte(str2)
//--stringת byte
// str = string(strByte[:])
/*---json str תmap
    jsonStr := this.GetString("data")
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==============json str תmap=======================")
		fmt.Println(dat)
		fmt.Println(dat["host"])
	}
*/
//----