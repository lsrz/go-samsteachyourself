package main

import (
	"fmt"
	"reflect" //reflect包实现了运行时反射
)

func main() {
	var a bool
	fmt.Println(a)

	a = true
	fmt.Println(a)

	var b int
	var c float32
	var d float64
	var e string

	var f [4]string
	f[0] = "a0"
	f[1] = "b1"
	f[2] = "c2"
	f[3] = "d3"

	var g = [4]int{1, 2, 3, 4}
	h := [10]int{1, 2, 3, 4, 5}

	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(f))

	//strconv.ParseInt()

	var i error
	fmt.Println(reflect.TypeOf(i))

}
