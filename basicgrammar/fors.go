/*
for 循环
*/
package basicgrammar

/**
go已package为最小单位 定义的变量如果没有被引用idea是会报错的
这是一个for
func name的大写是 public的属性
*/
import (
	"fmt"
)

//常量 const 访问权限也不一样 pkg都可以访问 使用
const b = "1213"

// break 中断整个循环 continue 中断当前循环 return 回调整个func
func Grammar() {
	//变量 只能在函数 Grammar 使用
	//变量命名规则 var 参数名称 类型 = 表达式
	var s, sep string
	//自动类型推断
	var zi = "12321"
	var a [3]string = [3]string{"a", "b", "c"}
	//简短变量声明 自动推导
	jianduan := "12321"
	for i := 0; i < len(a); i++ {
		s += sep + a[i]
		sep = ""
	}
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(zi)
	fmt.Println(zi)
	fmt.Println(jianduan)
	//获取变量zi的指针
	p := &zi
	fmt.Println(p)
	//获取指针的值
	fmt.Println(*p)
	//可以给指针指向的变量重新赋值
	*p = "0"
	fmt.Println(zi)
	//任何指针的零值是 nil 不理解没有测试出来
	fmt.Println(&zi == nil)
}
