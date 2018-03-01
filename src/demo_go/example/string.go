package example

import (
	"strings"
	"fmt"
)

//golang字符串去除空格和换行符
func replace_Blank(){
	str := "welcome to bai\ndu\n.com"
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println(str)
}

func main(){
	replace_Blank()
}