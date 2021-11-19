package main

import (
	"fmt"
	"time"
)

var b = 1

func main() {

	//if ... else if ... else ...
	if b == 1 {
		fmt.Printf("if b %v\n", b)
	} else if b == 2 {
		fmt.Printf("else if b %v\n", b)
	} else {
		fmt.Printf("else b %v\n", b)
	}
	// **没有括号,括号是多余的**

	//比较运算符 == ！= > >= < <=
	//算数运算符 + - * 、 %
	//逻辑运算符 ||(或) &&(与) !(非)

	// switch
	switch b {
	case 1:
		fmt.Println("switch One")
	case 2:
		fmt.Println("switch Two")
	default:
		fmt.Println("switch Defalut")
	}
	// **没有括号**
	// **没有break**

	//for
	i := 0
	for i < 3 {
		fmt.Printf("i=%v\n", i)
		i++
	}
	for j := 0; j < 3; j++ {
		fmt.Printf("j=%v\n", j)
	}
	//range
	k := [4]int{1, 2}
	for key, val := range k {
		fmt.Printf("key:%v, val:%v\n", key, val)
	}

	//defer:推迟,作用：执行清理,[dɪˈfɜː(r)][地缝]
	defer fmt.Println("I'm defer, i am run after the function completes，one") //最后执行
	defer fmt.Println("I'm defer, i am run after the function completes，two")
	defer fmt.Println("I'm defer, i am run after the function completes，thr") //最先执行
	fmt.Println("I'm not defer")
	// **先进后出**

	//
	t := time.NewTicker(time.Second * 1)
	for {
		<-t.C
		fmt.Println("ping")
		//time.Sleep(time.Second * 1)
	}
}
