package main

import (
	"fmt"
	"time"
)

/*
	不要通过共享内存来通信，而通过通信来共享内存
	Go语言使用通道在Goroutine之间收发消息，避免了使用共享内存
	通过收发消息，使得能够以推送方式协调并发事件。
*/
var c = make(chan string)

func SlowFunc2() {
	time.Sleep(time.Second * 2)
	//发送方
	c <- "Begin Sending, Can you Receive?"

	///可能会看不到他执行完哦, 所以应该最后一个动作是发送
	fmt.Println("SlowFunc2 Finished, Maybe Sometimes you can not see me!")
}
func main() {
	//c := make(chan string)
	//并发
	go SlowFunc2()

	//接收方(它将阻止函数返回，直到收到一条消息为止)
	msg := <-c
	fmt.Println(msg)

	fmt.Println("main finished!")
}
