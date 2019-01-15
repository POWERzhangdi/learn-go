/*
	interface
*/
package basicgrammar

import (
	"fmt"
	"reflect"
)

type IHuman struct {
	name  string
	age   int
	phone string
}

type IStudent struct {
	//匿名字段
	IHuman
	school string
	loan   float32
}

type IEmployee struct {
	//匿名字段
	IHuman
	company string
	money   float32
}

//printf 这里f 就是 format格式化输出语句
//IHuman的 ISayHi方法
func (h *IHuman) ISayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//IHuman的 Sing方法
func (h *IHuman) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//IHuman对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//IEmployee重载IHuman的ISayHi方法
func (e *IEmployee) ISayHi() {
	//此句可以分成多行
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//IStudent实现BorrowMoney方法
func (s *IStudent) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//IEmployee实现SpendSalary方法
func (e *IEmployee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

//定义interface
type Men interface {
	ISayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type YoungChap interface {
	ISayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	ISayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

//空interface

//定义a为空的interface
var a interface{}
var i int = 5

//这里是 没有使用 变量s idea自动识别错误 语法没有错 我先注释掉
//s := "Hello World"
//a = s
//空interface的优势 可以动态的传入不同的值 也可以回调 加入说回调是interface的时候
//a = i

//interface 存储的类型 方向知道存储的类型
//语法 value,ok = element.(T)
//go 嵌入interface 意思就是 interface A 所有的接口可以潜入到 interfaceB当中

//反射
//reflect 包做反射  1.先把空的interface转成reflect对象调用(reflect.Type || reflect.Value)
/**
t :=reflect.TypeOf(i) 获取到元数据，通过t获取类型定义里面的所有元素
v := reflect.ValueOf(i) 得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
*/

func IReflect() reflect.Type {
	var a = "2132"
	v := reflect.ValueOf(a)
	return v.Type()
}

//反射修改变量的值
func IEditReflect() reflect.Value {
	var a = "2132"
	//获取指针
	v := reflect.ValueOf(&a)
	//获取value
	e := v.Elem()
	//重新赋值
	e.SetString("12")
	return e
}
