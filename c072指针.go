package main

import (
	"fmt"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	//值引用
	//b := a
	//指针引用
	b := &a
	//b.Ice = false
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)

	//总结：
	//1.b是指针，使用指针的值加*
	//2.b是指针，打印内存地址不加&
	//3.&b应该代表的是变量b的内存地址
	//4.("#{a}\n")是IDE显示问题
}
