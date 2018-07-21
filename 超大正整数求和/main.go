package main

import (
	"fmt"
)

func main()  {
	c:
	a:=""
	b:=""
	e:=""
	fmt.Println("请输入两个正整数，每次输入完成，按回车确认")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(for1(a,b))
	a1:
	fmt.Println("输入c进行新一轮的计算;输入exit退出")
	fmt.Scan(&e)

	if e=="c" || e=="C"{
		goto c
	}else if e!="exit"{
		goto a1
	}

}
func for1(a,b string)string  {
	s:=""
	//判断长度
	if len(a)>len(b){
		swap:=a
		a=b
		b=swap
	}
	aa:=[]byte(a)
	bb:=[]byte(b)
	if len(bb)!=len(aa){
		for i:=range bb{
			if i+1==len(bb){
				bb[0]-=48
				break
			}
			if i >= len(aa){//单独运算
			    if bb[len(bb)-i-1]>57{
			    	bb[len(bb)-i-1]-=58
			    	bb[len(bb)-i-1-1]++
			    }else {
			    	bb[len(bb)-i-1]-=48
			    }

			}else {
				if bb[len(bb)-i-1]+aa[len(aa)-i-1]-96>9{
					bb[len(bb)-1-i-1]++
					bb[len(bb)-i-1]=bb[len(bb)-i-1]+aa[len(aa)-i-1]-106
				}else {
					bb[len(bb)-i-1]=bb[len(bb)-i-1]+aa[len(aa)-i-1]-96
				}
			}

		}
	}else {
		for i:=range bb{
			if i+1==len(bb){
				bb[len(bb)-i-1]=bb[len(bb)-i-1]+aa[len(aa)-i-1]-96
				break
			}

			if bb[len(bb)-i-1]+aa[len(aa)-i-1]-96>9{
				bb[len(bb)-1+i-1]++
				bb[len(bb)-i-1]=bb[len(bb)-i-1]+aa[len(aa)-i-1]-106
			}else {
				bb[len(bb)-i-1]=bb[len(bb)-i-1]+aa[len(aa)-i-1]-96
			}

		}
	}

	for _,v:=range bb{
		vv:=fmt.Sprintf("%v",v)
		s+=vv
	}
	return s
}