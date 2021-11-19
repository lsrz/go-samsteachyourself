package main

import "fmt"

//我习惯了将常量名全大写，如CONSTANT，在Go语言中，也应这样做吗？
//在Go语言中，常量与其他数据元素没什么不同，因此不应将其全大写。
const hello = "hello, const"

func main() {

	fmt.Println(hello)
}
