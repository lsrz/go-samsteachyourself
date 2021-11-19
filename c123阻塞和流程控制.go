package main

import (
	"fmt"
	"time"
)

/*
	Goroutine 是非阻塞的
	因此如果程序要阻塞，以接收大量的消息或不断地重复某个过程，必须使用其他流程控制技术
*/

var msg = make(chan string)

func sfc() {
	t := time.NewTicker(time.Second * 1)
	for {
		msg <- "ping"
		//time.Sleep(time.Second*1)
		<-t.C
	}
}

func main() {
	go sfc()
	for {
		message := <-msg
		fmt.Println(message)
	}
}
