package main

import (
	"fmt"
	"encoding/json"
	//"os/exec"
	"bufio"
	"io"
	"os"
	"strconv"
)
func main1(){
	fmt.Sprintf("Hello, %s!", "dddd")
	fmt.Println("Hello, World!")
}

func array_append(){ //demo
	slice:=[]int{}
	for i:=0;i<20;i++{
		slice=append(slice,i)
		fmt.Println(cap(slice))
		fmt.Println(slice)
	}
}
func file_readline(){
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行 //ReadString（）的函数说明, err！=nil，也会把已经读到的数据返回出来的。
		fmt.Println(line)
		if err != nil || io.EOF == err {
			break
		}
	}
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

func jsonParse(f interface{}) int{
	m := f.(map[string]interface{})
	for k, v := range m {
		if (k==""){

		}
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
	return 0
}

type InfoRequest struct {
	Rd string
	Time string
	Key string
	From string
	Data interface{}
}
func checkRequest(info InfoRequest) bool{
	if (info.Rd=="test") {
		return true
	}
	return false
}
type InfoServer struct {
	Success int
	Time string
	Msg string
	Data interface{}
}

func main() {
	var infoServer *InfoServer = &InfoServer{
		Success: 1,
		Time: "",
		Msg: "",
	}
	infoServer.Time=""
	//----
	var infoRequest InfoRequest
	//jsonStr := []byte(`{"Name":"aree", "Age":18,"From": [ "SZ", "GD" ],"data":[{"a":"aa","b":null},{"c":[]},{"list":["dd",1,"650827..."]}]}`)
	//jsonStr := []byte(`[{"Name":"aree", "Age":18,"From": [ "SZ", "GD" ],"data":[{"a":"aa","b":null},{"c":[]},{"list":["dd",1,"ff"]}]}]`)
	//jsonStr := []byte(`{"succ":0,"rd":"st","key","001001","data":{"a":"aaa"}}`)
	jsonStr := []byte(`{"rd":"test","cmd":"001001","data":{"a":"aaa"},"key":"001001","time":"2018-03-01 11:38:54"}`)
		var f interface{}
		err := json.Unmarshal(jsonStr, &f)
		if err != nil {
			fmt.Println(err)
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
				infoRequest.Rd=v.(string)
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
			fmt.Println("Error:",infoRequest.Rd)
		}


}


//cmd := exec.Command("/bin/sh","-c",`curl -d sysid=xxx -d pwd=xxx -d tel=xxx -d warn=xxx http://xxx`)
// 我们知道interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：
//bool 代表 JSON booleans,
//float64 代表 JSON numbers,
//string 代表 JSON strings,
//nil 代表 JSON null.