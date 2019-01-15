/**
go关键字多线程
runtime 中处理 goroutine的函数
Goexit 退出当前执行的goroutine defer函数还会继续调用
Gosched 让出当前 goroutine的执行权限，调度安排其他等待的任务运行，并在下次某个时候从该位置恢复执行
NumCPU 返回cup核数量
NumGoroutine 返回正在执行和排队的任务总数
GOMAXPROCS 设置可以并行计算的CPU核数的最大值，并返回之前的值
*/
package basicgrammar

import (
	"fmt"
	"runtime"
	"time"
)

/*
go 关键字只是做到了并发 要想多核处理还要 显示的初始化调用runtime.GOMAXPROCS(n)
告知调度器同时使用多个线程
*/

func GoSay(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

/*
channels goroutine运行在相同的内存空间 要访问共享的内存必须做好同步
goroutine 之间如何通信 就采用channel通信
发送接收都是特定的类型 channel类型 必须采用mark创建channel
c := make(chan int)
ch <-v 发送v到channel ch
v:= <-ch 从ch中接收数据 赋值给v
*/

func Sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	//发送total 到c通道
	c <- total
}

/*
Buffered Channels go可以指定channel缓存的大小 就是可以存储多少个元素
ch := make (chan bool,4) 指定4个元素 在前4个之前channel不会阻塞可以写入
当到达4之后 第5个元素不可写入阻塞
ch := make (chan type ,value)
注意点：value = 0 无缓冲阻塞读写的
*/
func TbufferChannels() {
	// chan 的值 修改为1报如下的错误: fatal error: all goroutines are asleep - deadlock!
	c := make(chan int, 1)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/*
通过 Range Close 操作多次channel
*/

func Fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/*
多个 channel的时候采用 Select选择 这个跟netty很类似
*/

func SelectFibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			//默认执行不阻塞的channel
		}
	}
}

//超时

func TimeOut() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
