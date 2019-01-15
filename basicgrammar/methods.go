/*
method
*/
package basicgrammar

import "fmt"

type Rectangle struct {
	width, height float64
}

//method 语法 func (r ReceiverType) funcName(parameters) (results)
func (r Rectangle) area() float64 {
	return r.height * r.width
}

//传统形式
func Area(r Rectangle) float64 {
	return r.height * r.width
}

//method重写
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

//Human定义method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}
