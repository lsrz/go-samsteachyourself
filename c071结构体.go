package main

import (
	"fmt"
	"reflect"
)

//为了单个变量中存储众多类型不同的数据宇段

//type 指定一种新类型
//Person 新类型名称
//struct 数据类型为结构体

//可将结构体视为模板
//根据Go语言约定，结构体及其数据宇段都可能被导出，也可能不导出。如果一个标识符的首字母是大写的，那么它将被导出；否则不会导出。
//要导出结构体及其字段，结构体及其宇段的名称都必须以大写字母打头。
type Person struct {
	Name string
	Age  int8
	male bool
}

//嵌套
type Class struct {
	Name string
	Nums int8
}
type Grade struct {
	Name  string
	Nums  int
	Class Class
}

func funshow(grade string, class string) Grade {
	a := Grade{
		Nums: 10,
	}
	a.Name = grade
	a.Class.Name = class

	return a
}

func main() {
	p := Person{
		Name: "lucy",
		Age:  18,
		male: false,
	}
	fmt.Println(p)

	//名词：点表示法（结构体变量名、圆点和要访问的数据宇段）
	fmt.Println("Name:", p.Name, "Age:", p.Name)

	//创建结构体
	var px = Person{}
	//打印结构体 %+v
	fmt.Printf("%+v\n", px)
	//零值
	//{Name: Age:0 male:false}
	//赋值（点表示法）
	px.Name = "lily"
	fmt.Printf("%+v\n", px)

	//创建结构体
	py := new(Person)
	pz := Person{"hanmm", 17, false} //不推荐
	fmt.Printf("%+v %+v\n", *py, pz)
	//Think about why there is a star in front of py?

	gc := Grade{
		Name: "一年级",
		Nums: 10,
		Class: Class{
			Name: "一年一班",
			Nums: 41,
		},
	}
	fmt.Printf("%+v\n", gc)

	//***理解type***//
	//数组
	var gc1 [2]Grade
	gc1[0] = gc
	fmt.Printf("%+v\n", gc1)
	//切片
	var gc2 = make([]Grade, 1)
	gc2[0] = gc
	fmt.Printf("%+v\n", gc2)
	//函数
	fshow := funshow("二年级", "二年一班")
	fmt.Printf("%+v\n", fshow)

	fmt.Println(reflect.TypeOf(gc))
	fmt.Println(reflect.TypeOf(gc1))
	fmt.Println(reflect.TypeOf(gc2))
	fmt.Println(reflect.TypeOf(fshow))

	if gc == fshow {
		fmt.Println("equal")
	} else {
		fmt.Print("not equal:")
		fmt.Printf("%+v,%+v\n", gc, fshow)
	}

	//结构体切片表格，参看 doctest/example02_test.go
}
