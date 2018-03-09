package main

import (
	"net"
	"log"
	"io"
	"fmt"
	"bufio"
	"strings"
	"os"
	"path/filepath"
	"io/ioutil"
)

var (
	cmd       string
	file_name string
)

func main() {
	addr := "0.0.0.0:81" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	conn, err := listener.Accept() //用conn接收链接
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("欢迎来到迷你FTP服务器！"))
	r := bufio.NewReader(conn) //将这个链接（connection）包装以下。将conn的内容都放入r中，但是没有进行读取，让步我们一会对其进行操作。
	for {
		line, err := r.ReadString('\n') //将r的内容也就是conn的数据按照换行符进行读取。
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(line)
		line = strings.TrimSpace(line) //去掉换行符。
		fmt.Println(len(strings.Fields(line)))
		if len(line) == 0 { //为了让客户端长时间和服务器通话。
			continue
		}
		cmd = strings.Fields(line)[0]
		if len(strings.Fields(line)) > 1 {
			file_name = strings.Fields(line)[1] //需要获取服务器的文件
		}
		pwd, err := os.Getwd()
		if err != nil {
			panic("获取路径出错了！")
		}
		file_name = filepath.Join(pwd, file_name)
		fmt.Println(file_name)
		switch  cmd {
		case "GET", "get":
			f, err := os.Open(file_name) //打开文件的内容。
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()
			buf, err := ioutil.ReadAll(f)
			if err != nil {
				log.Print(err)
				return
			}
			conn.Write(buf)
		case "PUSH", "push":
			fmt.Println("上传文件的语句")
			conn.Write([]byte("上传文件的命令\n"))

		case "EXIT", "exit":
			//conn.Close()
			return
		default:
			fmt.Println("您输入的命令无效！")
			conn.Write([]byte("您输入的指令有问题!\n"))
		}
	}
}
