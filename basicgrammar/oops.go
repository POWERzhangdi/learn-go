/*
面向对象
*/
package basicgrammar

/**
普通的计算方式
跨包引用 属性值也要大写
大写即public
*/
type Chang struct {
	Width, Height float64
}

/**
Azhang 可以实现计算但是相当于一个 function(method)
不是 Chang 应有的属性 不属于面向对象
*/
func Ajisuan(r Chang) float64 {
	return r.Width * r.Height
}

type Yuan struct {
	Width, Height float64
}

/*
	go method 让它成为面向对象
	receiverType 依附着依附在谁的身上
	method语法 func (r receiverType) funcName(parameters)(result)
*/

//此时 Yuan 对象身上就包含了 MethodJisuan
func (y Yuan) MethodJisuan(yuan Yuan) float64 {
	return yuan.Height * yuan.Width
}

func (c Chang) MethodJisuan(chang Chang) float64 {
	return chang.Height * chang.Width
}

//自定义类 猜测 string 是 key的类型，int是value的类型
type customClass map[string]int

func DeclareCustomClass() customClass {
	custom := customClass{
		"first":  1,
		"second": 2,
		"third":  3,
	}
	return custom
}
