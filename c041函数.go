package main

import (
	"fmt"
	"strconv"
)

// 1.声明返回值类型
// 2.func行被称作函数签名
func addUP(x int, y int) int {
	return x + y
}

// 返回多个值
func someFun() (int, string) {
	return 2, "aaa"
}

// 不定参数...
func sumNumbers(nums ...int) int {

	total := 0
	// 1. _ 是个特殊的标识符，代表某个值我们不关心，接收了但不使用
	// 2. range在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
	//      for k, v := range kvs {
	//	      fmt.Printf("%s -> %s\n", k, v)
	//      }
	for _, num := range nums {
		total += num
	}
	return total
}

// 具名返回值
// x y 称为具名变量
// return 称为"裸"return （naked return）
func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}

//递归函数
func eat(every int, eaten int) string {
	eaten = every + eaten

	if eaten >= 5 {
		return "i'm full, i have eaten " + strconv.Itoa(eaten) + "\n"
	}
	return eat(every, eaten) + "i'm hungry, i have eaten " + strconv.Itoa(eaten) + "\n"
}

//将函数作为参数传递
//f func() string也是一个函数签名
func AnotherFunc(f func() string) string {
	return f()
}

func main() {
	num, val := someFun()
	fmt.Printf("the num is: %v, the string is: %v\n", num, val)

	total := sumNumbers(1, 2, 3, 4)
	fmt.Printf("total is: %v\n", total)

	fmt.Println(sayHi())
	fmt.Println(eat(1, 0))

	a := func() string {
		return "a"
	}
	fmt.Println(AnotherFunc(a))
}
