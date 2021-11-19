package main

import (
	"errors"
	"fmt"
	"os"
)

//Go语言处理错误的方式很有趣---将错误作为一种类型，这意味着可将错误传递给函数和方法。
//Go语言将错误作为返回值,不支持传统的try-catch-finally控制结构
func main() {

	//file, err := ioutil.ReadFile("foo.txt")
	file, err := os.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", file)
	}

	//创建错误
	err = errors.New("Something went wrong")
	if err != nil {
		fmt.Println(err)
	}

	name, role := "Richard Jupp", "Drummer"
	err = fmt.Errorf("The %v %v quit", role, name)
	if err != nil {
		fmt.Println(err)
	}
	//Errorf返回的是字符串，可以赋值。
	str := fmt.Sprintf("Sprintf %v %v test", role, name)
	fmt.Println(str)

	//返回错误 error is a type too && return nil
	//1. error 是一种类型（type）
	//2. 直接返回nil的写法
	//3. 多结果返回
	n, err := Half(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}

	//panic [ˈpænɪk] 惊恐;恐慌;
	//程序立即停止执行
	panic("Oh no. I can do no more. Goodbye .")
}

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v ", numberToHalf)
	} else {
		return numberToHalf / 2, nil
	}
}
