package main

import (
	"fmt"
)

//Type can be omitted(省略)
var s string = "s"

//s = "b" go不能在函数外赋值

func main() {

	//全部声明方式
	var s1 string = "s1"
	var s2 = "s2"
	var s3 string
	s3 = "s3"
	s4 := "s4" //短变量声明

	//快捷变量声明
	var s5, s6 string = "s5", "s6"
	var c, python, java bool

	//简短变量声明，只能在函数中使用
	a1, b1, c1, d1, f1 := 3, 4, 22, "iii", true
	//file, err := os.ReadFile("foo.txt")

	// var声明一个变量
	// 变量名
	// 变量类型

	// 变量不能重复声明
	// 声明变量后如果不指定值，go将指定默认值，称为零值，根据类型不同一般为0，false，""
	// go要求不能在函数外部使用简短变量声明，简短变量声明是最常用的方式

	// var s8 这是错误的
	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer

	// 这是错误的（undefined: a），必须先声明后使用
	//fmt.Printf("hello %v", a)
	//a := "a"

	// nil：指向oc中对象的空指针。
	// Nil：指向oc中类的空指针。
	// NULL：指向其他类型的空指针，如一个c类型的内存指针。

	fmt.Println(s, s1, s2, s3, s4, s5, s6, c, python, java)
	fmt.Println(a1, b1, c1, d1, f1)
}
