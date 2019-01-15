/*
指针
*/
package basicgrammar

import "fmt"

//函数是一个变量
var p = f()

func f() *int {
	v := 1
	fmt.Println(&v)
	return &v
}

func incr(p *int) int {
	//非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	*p++
	//上面的等价于 *p = *p + 1
	return *p
}

func TestPoint() {
	a := incr(p)
	fmt.Println(a)
}
