package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main()  {
	str:=""
	fmt.Println("请输入要连接的服务端ip与端口,参考格式:127.0.0.1:8080")
	fmt.Scan(&str)
	conn,err:=net.Dial("tcp",strings.TrimSpace(str))
	pe(err)
	fmt.Println("与",str,"连接成功")
	fmt.Println("请输入要连接的客户端的ip与端口,参考格式:127.0.0.1:8080")
	send("client",conn)
a:
	go func() {
		fmt.Println("请输入要发送的消息;想退出请输入:exit;")
		send("msg",conn)
	}()
	buf:=make([]byte,1024)
	for{
		count,err:=conn.Read(buf)
		if nil!=err{fmt.Sprint("连接不上服务器");return}
		if count>0{
			fmt.Println(string(buf[0:count]))
			goto a
		}
	}
}

func send(str string,conn net.Conn)  {
	switch str {
	case "client":
	a:
		input,_,_:=bufio.NewReader(os.Stdin).ReadLine()
		input=[]byte(strings.TrimSpace(string(input)))
		conn.Write(input)
		//判断对方是否在线
		buf:=make([]byte,30)
		count,err:=conn.Read(buf)
		if nil!=err{fmt.Sprint("连接不上服务器");return}
		if count>0{
			if string(buf[0:count])=="1"{
				fmt.Println(string(input),"客户端未连接服务器,请重新输入,连接别的客户端")
				goto a
			}else {fmt.Println("连接",string(input),"成功")}
		}
	case "msg":
		input,_,_:=bufio.NewReader(os.Stdin).ReadLine()
		if strings.ToLower(strings.TrimSpace(string(input)))=="exit"{os.Exit(0)}
		conn.Write(input)

	}
}
func pe(e error)  {
	if nil!=e{panic(e)}
}