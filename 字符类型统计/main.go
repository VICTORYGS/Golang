package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	a:
	fmt.Println("请输入字符，分别统计字母、空格、数字、基本汉字体（Unicode:19968-40869）、其他字符的个数")

	read:=bufio.NewReader(os.Stdin)

	res,_,err:=read.ReadLine()

	if err!=nil{
		fmt.Println("read from console err:",err)
		return
	}

	letter:=0
	space:=0
	number:=0
	other:=0
	chinese:=0

	 for _,v:=range string(res){
		 switch {
		 case (v>='A'&&v<='Z')||(v>='a'&&v<='z'):
		 	letter++
		 case v<'9'&&v>'0':
		 	number++
		 case v==' ':
		 	space++
		 case v>'一'&&v<'龥':
			 chinese++
		 default:
			 other++
		 }

	 }

	fmt.Println("\n汉字:",chinese,"\n字母:",letter,"\n数字:",number,"\n空格:",space,"\n其他:",other)

	b:
	fmt.Println("\n\n输入exit 退出\n输入c 重新统计")

	directive:=""

	fmt.Scanf("%v\n",&directive)

	if directive=="c"||directive=="C"{
		goto a
	}else if directive!="exit"{
		goto b
	}


}
