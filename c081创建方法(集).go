package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

func summary(m *Movie) string {
	return m.Name + ":" + strconv.FormatFloat(m.Rating, 'f', 1, 64)
}

//方法比函数多一个接收者
func (m *Movie) summary() string {
	return m.Name + ":" + strconv.FormatFloat(m.Rating, 'f', 1, 64)
}

func main() {

	m := Movie{
		Name:   "1st blood",
		Rating: 32.0,
	}

	//函数
	fmt.Println(summary(&m))
	//方法
	fmt.Println(m.summary())

	//函数和方法的区别联系？WTF
	//https://www.cnblogs.com/zuoyang/p/15175276.html

}

//方法集：很多方法
//...
