package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func multiply(d int) int {
	if d==1{return 1}
	return multiply(d-1)*d
}
func main()  {
	a:

	fmt.Println("请输入一个不大于25的正整数，用来计算阶乘")
	d:=0
	read:=bufio.NewReader(os.Stdin)
	res,_,err:=read.ReadLine()
	if err==nil{
		d,_=strconv.Atoi(string(res))
	}

	fmt.Printf("!%v = %v",d,multiply(d))

b:
	fmt.Println("\n\n输入exit 退出\n输入c 重新计算")

	directive:=""

	fmt.Scanf("%v\n",&directive)

	if directive=="c"||directive=="C"{
		goto a
	}else if directive!="exit"{
		goto b
	}

}
