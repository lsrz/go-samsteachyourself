package main

import "fmt"

func sayHello(s string) string {
	return "Hello," + s
}

func addition(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello," + "Andy1")
	fmt.Println(sayHello("Andy2"))
	fmt.Println(addition(2, 3))
}
