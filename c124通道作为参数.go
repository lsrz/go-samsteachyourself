package main

import (
	"fmt"
)

//发送方
func SendOnly(m chan<- string) {
	m <- "Hello"
}

//接收方
func ReceiveOnly(m <-chan string) string {
	msg := <-m
	return msg
}

//发送and接收
func SendAndReceive(m chan string) {
	msg := <-m
	fmt.Println(msg)
	m <- "Hello"
}

func main() {
	var msg124 = make(chan string)
	go SendOnly(msg124)
	m := ReceiveOnly(msg124)
	fmt.Println(m)
}
