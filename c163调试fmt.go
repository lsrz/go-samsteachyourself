package main

import "fmt"

func main() {
	//fmt
	s := "Hello World"
	t := "Goodbye, Cruel World "
	fmt.Printf("s is %v, t is %v\n", s, t)

	//
	var floor3 = []struct {
		name string
		age  int
	}{
		{"lilei", 22},
		{"hanmm", 22},
		{"lucy", 22},
		{"lily", 22},
	}

	for _, item := range floor3 {
		fmt.Printf("%v\n", item)
		fmt.Printf("%+v\n", item)
		fmt.Printf("%v:%v\n", item.name, item.age)
	}

}
