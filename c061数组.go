package main

import "fmt"

var cheeses [2]string

func main() {

	cheeses[0] = "a"
	cheeses[1] = "b c"

	//打印没有逗号分割 [a b c]
	fmt.Println(cheeses)

	//创建长度为3的int数组
	array1 := [3]int{10, 20, 30}
	fmt.Println(array1)
	array2 := [4]byte{'b', 'y', 't', 'e'}
	fmt.Println(array2)
	array3 := [1]string{"byte"}
	fmt.Println(array3)
}
