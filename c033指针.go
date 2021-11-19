package main

import "fmt"

func main() {
	g := "hi,go"
	o := "hi,go"

	//输出内存地址 0xc00003a230
	fmt.Println(&g)
	fmt.Println(&o)

	a := 2
	fmt.Println(&a)
	usePointerAdd(&a)
	usePointerVal(&a)

}

//声明指针，在类型前加星号，如：*int
func usePointerAdd(x *int) {
	//打印内存地址
	fmt.Println(x)
	return
}

//使用指针的值，在变量前加星号，如：*x
func usePointerVal(x *int) {
	//打印变量值
	fmt.Println(*x)
	return
}

//将变量传递给函数时，会分配新的内存并将变量的值复制到其中。
//这样将产生两个变量的实例，他们位于不同的内存单元中。
//一般而言这不可取。
//因为这将占用更多的内存，同时由于存在变量的多个副本，很容易引入bug
//所以go引入指针。
