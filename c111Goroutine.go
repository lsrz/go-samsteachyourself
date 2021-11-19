package main

import (
	"fmt"
	"time"
)

/*
routine  [ruːˈtiːn]

顺序执行：balabala
并发：在一台处理器上同时处理多个任务
并行：在多台处理器上同时处理多个任务

阻塞：直到响应返回为止
非阻塞：balabala

go SlowFunc()
在调用函数SlowFunc后立即执行main函数中的第二行代码,不会阻塞程序中其他代码行的执行。
程序将在Goroutine返回前退出,这样才符合非阻塞执行的理念。

Goroutine与线程
Goroutine是建立在线程之上的轻量级的抽象。它允许我们以非常低的代价在同一个地址空间中并行地执行多个函数或者方法。
相比于线程，它的创建和销毁的代价要小很多，并且它的调度是独立于线程的。
Goroutine所需要的内存通常只有2kb，而线程则需要1Mb（500倍）。
当一个goroutine被阻塞（比如等待IO），golang的scheduler会调度其它可以执行的goroutine运行。
*/

func SlowFunc() {
	/*
		首先：time.sleep单位为:1ns (纳秒)
		转换单位：
			  1纳秒 =1000皮秒
			  1纳秒 =0.001 微秒
			  1纳秒 =0.000 001毫秒
			  1纳秒 =0.000 000 001秒
		time.Second = 1s
	*/
	time.Sleep(time.Second * 2)
	fmt.Println("SlowFunc finished!")
}
func main() {
	//非阻塞，程序将在Goroutine返回前退出！
	go SlowFunc()

	fmt.Println("main finished!")
	//避免不等待并发完成直接结束
	time.Sleep(time.Second * 3)
}
