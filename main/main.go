package main

import (
	"fmt"
	"learn/basicgrammar"
)

func main() {
	//basicgrammar.Grammar()
	//basicgrammar.TestPoint()
	//basicgrammar.Person()
	//mark := basicgrammar.Student{basicgrammar.Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	//sam := basicgrammar.Employee{basicgrammar.Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	//mark.SayHi()
	//sam.SayHi()
	//计算一个长方形的面积
	//c1 :=basicgrammar.Chang{Width:12,Height:2}
	//fmt.Println(basicgrammar.Ajisuan(c1))
	//fmt.Println(basicgrammar.IReflect())
	//fmt.Println(basicgrammar.IEditReflect())
	//go basicgrammar.GoSay("World")
	//basicgrammar.GoSay("Hello")
	//a:=[]int{7, 2, 8, -9, 4, 0}
	//c :=make(chan int)
	//go basicgrammar.Sum(a[:len(a)/2],c)
	//go basicgrammar.Sum(a[len(a)/2:], c)
	////读取操作会被阻塞直到 有数据接收 发送方也会被阻塞 直到数据被读出 不用显示申明lock 本身就阻塞
	//x ,y := <-c,<-c //从c通道当中获取值赋值
	//fmt.Println(x,y,x + y)
	//basicgrammar.TbufferChannels()
	//c := make(chan  int,10)
	//go basicgrammar.Fibonacci(cap(c),c)
	//for i:=range c {
	//	fmt.Println(i)
	//}
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	basicgrammar.SelectFibonacci(c, quit)
}
