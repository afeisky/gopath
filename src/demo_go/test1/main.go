package main

import (
	"fmt"
	"encoding/json"
	"time"
)

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

func main() {
	fmt.Println("App1DataController ->")

	jsonStr := `{"Name":"app1","Time":"2018-03-08","Ver":"v1.0","From":"addr","Data":"aa"}`
	fmt.Println(jsonStr)
	//parse receiver data
	var rev InfoReceiv
	err := json.Unmarshal([]byte(jsonStr), &rev)
	if err != nil {
		fmt.Println(err)
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
}
func main2() {
	fmt.Println("App1DataController ->")

	jsonStr :=`{"Name":"app1","Time":"2018-03-08","Ver":"v1.0","K1":"key1","K2":"key2","From":"addr","Data":"aa"}`
	fmt.Println(jsonStr)
	jsonData := []byte(jsonStr)
	var f interface{}
	err := json.Unmarshal(jsonData, &f)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ctime string
	var rd string
	var data interface{}
	m := f.(map[string]interface{})
	var rev InfoReceiv
	for k, v := range m {
		switch k {
		case "Name":
			rd=v.(string)
			rev.Name=v.(string)
		case "Time":
			ctime=v.(string)
			rev.Time=v.(string)
		case "Data":
			data=v
			rev.Data=v
		default:
			fmt.Println(k, "[type?]",v)
		}
	}
	fmt.Println("rd",rd)
	fmt.Println("ctime",ctime)
	fmt.Println("data",data)

	var send *InfoSend = &InfoSend{
		Success: 0,
		Time: "",
		Msg: "success",
	}
	send.Time=rev.Time
	//---
	js, _ := json.Marshal(send)
	fmt.Println(string(js))//{"success":1,"time":""}

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
