package main

import (
	"fmt"
	"time"
)

// 使用退出通道
// stp := make(chan bool)
// stp <- true

func main() {

	msg := make(chan string)
	stp := make(chan bool)

	go func() {
		t := time.NewTicker(time.Second * 1)
		for {
			<-t.C
			msg <- "send after 1 sec"
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("send stop")
		stp <- true
	}()

	//for i := 1; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 3 {
	//		return
	//	}
	//}
	for {
		select {
		case <-stp:
			//fmt.Println("stop in select")
			return
		case ret := <-msg:
			fmt.Println(ret)
		}
	}
}
