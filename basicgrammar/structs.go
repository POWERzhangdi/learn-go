/*
自定义对象
*/
package basicgrammar

import "fmt"

//type 声明一个类型 可以看作java的class name ，age是 person的属性 go语言当中需要  struct关键字
type person struct {
	name string
	age  int
}

type person1 struct {
	//这里引用是是 person对象 那么 person1 对象的 属性 name 就是person对象
	person person
	age    int
}

//属性赋值

func Person() {
	//第一种初始化
	var P person
	P.name = "张迪"
	P.age = 12
	fmt.Println(P.name)
	//第二种初始化 按照属性顺序执行
	P1 := person{"Tom", 25}
	fmt.Println(P1.name)
	//第三种初始化 可以按照指定顺序执行 field.value的形式
	P2 := person{name: "LiLi", age: 24}
	fmt.Println(P2.name)
	//为何age不能写出来？
	Person1 := person1{person{"TOM", 12}, 12}
	fmt.Println(Person1.person)
}
