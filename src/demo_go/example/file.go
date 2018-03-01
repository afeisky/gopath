package example

import (
	"os"
	"bufio"
	"io"
	"strings"
	"fmt"
	"os/exec"
	"container/list"
)

func readLine(){
	lines := list.New()

	f, err := os.Open("Y:\\a3a1\\11.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行 //ReadString（）的函数说明, err！=nil，也会把已经读到的数据返回出来的。
		//fmt.Println(line)
		lines.PushBack(line)
		if err != nil || io.EOF == err {
			break
		}
	}

	i := 0
	// 遍历
	for v := lines.Front(); v != nil; v = v.Next() {
		oldpath:="/disk1/a3a1/"+v.Value.(string)
		//newpath:="Y:/b/"+v.Value.(string)
		cmdline:=oldpath+" /disk1/b/"
		cmdline = strings.Replace(cmdline, "\n", "", -1)
		fmt.Println("cp ", cmdline)
		exec.Command("cp "+cmdline,"-r",``)
		i++
		if (i > 500) {
			break
		}
	}
}

func main(){

}