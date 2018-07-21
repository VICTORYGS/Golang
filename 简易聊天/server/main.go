package main

import (
	"net"
	"fmt"
	"strings"
)
var c =make(map[string]*struct{online bool;conn net.Conn;to string})
var ch=make(chan int,1)
func main()  {
	str:=""
	fmt.Println("请输入要监听的ip与端口,参考格式:127.0.0.1:8080")
	fmt.Scan(&str)
	dial,err:=net.Listen("tcp",strings.TrimSpace(str))
	pe(err)
	fmt.Println("server is waiting")
	for  {
		conn,err:=dial.Accept()
		pe(err)
		fmt.Printf("\n%v已连接\n",conn.RemoteAddr())
		ch<-1
		c[fmt.Sprintf("%v",conn.RemoteAddr())]= &struct{online bool;conn net.Conn;to string}{online:true,conn:conn}
		<-ch
		go process(conn)//conn是指针类型
	}
}
func process(conn net.Conn)  {
	buf:=make([]byte,2024)
	for i:=0;;i++{
		count,err:=conn.Read(buf)//不断读取
		if nil!=err{
			ch<-1
			c[fmt.Sprintf("%v",conn.RemoteAddr())]=&struct{online bool;conn net.Conn;to string}{online:false}
			<-ch
			//出错时说明连接已经断开
			fmt.Printf("\n%v已断开\n",conn.RemoteAddr())
			return
		}
		if count>0{
			if i==0{
				if c[string(buf[0:count])].online{
					ch<-1
					c[fmt.Sprintf("%v",conn.RemoteAddr())].to=string(buf[0:count])
					<-ch
					conn.Write([]byte("0"))
				} else {
					conn.Write([]byte("1"))
					i=0
				}
			}else {
				fmt.Printf("\n%v发来:%v 寄给->%v\n",conn.RemoteAddr(),string(buf[0:count]),(c[fmt.Sprintf("%v",conn.RemoteAddr())]).to)
				(c[(c[fmt.Sprintf("%v",conn.RemoteAddr())]).to]).conn.Write([]byte(fmt.Sprintf("%v发来：%v",conn.RemoteAddr(),string(buf[0:count]))))
				conn.Write([]byte("发送成功"))
			}
		}
	}
}
func pe(e error)  {
	if nil!=e{panic(e)}
}