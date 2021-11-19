package main

import (
	"fmt"
	"time"
)

//收到一条消息后， select 语句将不再阻塞。

var re1 = make(chan string)
var re2 = make(chan string)

func SendRet1() {
	time.Sleep(time.Second * 5)
	re1 <- "SendRet1"
}
func SendRet2() {
	time.Sleep(time.Second * 1)
	re2 <- "SendRet2"
}

func main() {

	go SendRet1()
	go SendRet2()

	//msg1 := <-re1
	//msg2 := <-re2

	select {
	//case msg1   			 // Error1
	//case <-re1:   		 // Yes
	//	fmt.Println(<-re1)   // Error2
	case msg1 := <-re1:      // 打印消息只有一种方法
		fmt.Println(msg1)
	case msg2 := <-re2:
		fmt.Println(msg2)
	case <-time.After(500 * time.Microsecond):
		fmt.Println("no messages received. giving up.")
	}
}
