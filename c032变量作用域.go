package main

import "fmt"

var outvar = "outvar"

func main() {
	fmt.Printf("hello %v\n", outvar)

	invar := true
	if invar {
		fmt.Printf("hello %v\n", invar)
	}

	// go 使用基于块的词法作用域
	// 一对{}标识一个块
	// 大括号内的大括号定义一个新块 - 内部块
	// 外部块不可以访问内部块的变量，内部块可以访问外部块的变量
	// go认为文件是一个块，第一级大括号外声明的变量可以在所有块访问

	sum := 0
	for i := 0; i <= 5; i++ {

		// 内部块测试
		innervar := true
		fmt.Println(innervar)

		sum += i
	}
	fmt.Println(sum)

	// 这是错误的(undefined: innervar)
	// fmt.Println(innervar)
}
