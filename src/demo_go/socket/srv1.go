package main

import (
	"net"
	"log"
	"time"
)

func main() {
	addr := "0.0.0.0:82"                     //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr) //使用协议是tcp，监听的地址是addr
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //关闭监听的端口
	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("tcl server\n")) //通过conn的wirte方法将这些数据返回给客户端。
		conn.Write([]byte("hello Golang\n"))
		time.Sleep(time.Minute) //在结束这个链接之前需要睡一分钟在结束当前循环。
		conn.Close()            //与客户端断开连接。
	}
}
