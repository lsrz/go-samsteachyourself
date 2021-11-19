package main

import "fmt"

//https://blog.csdn.net/tomatomas/article/details/95880842

func main() {

	fmt.Println("Type must be a pointer, channel, func, interface, map, or slice type")
	fmt.Println("nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值，也就是预定义好的一个变量")

	type A interface{}
	type B struct{}
	var a *A
	var b *B
	var c *int

	if a == nil {
		fmt.Println("a is nil")
	}
	if b == nil {
		fmt.Println("a is nil")
	}
	if c == nil {
		fmt.Println("c is nil")
	}

	var slice []string
	//var slice = make([]string, 3)
	if slice == nil {
		fmt.Println("slice is nil")
	}

}
