package main

import (
	"bytes"
	"fmt"
	"strings"
)

//Go语言中的字符串实际上是只读的字节切片
func main() {

	//创建字符串字面量
	s := "I am an interpreted string literal "
	fmt.Println(s)

	//使用rune字面量
	s = "After a backslash , certain single character escapes represent special values\nn is a line feed or new line \n\tt is a tab"
	fmt.Println(s)

	//原始字符串字面量
	s = `After a backslash , certain single character escapes represent special values
n is a line feed or new line 
	t is a tab`
	fmt.Println(s)

	//拼接字符串
	s = " Oh sweet ignition " + " be my fuse "
	s += " \nHear me screamin ’ ? "
	fmt.Println(s)

	//使用缓冲区拼接字符串(效率最高)
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("好")
	}
	//fmt.Println(buffer)
	fmt.Println(buffer.String())

	//字符串是切片
	s = "hello"
	fmt.Printf("字面量：%q \n", s[0]) //单引号围绕的字符字面量
	fmt.Printf("字节(byte)：%v \n", s[0])
	fmt.Printf("二进制：%b \n", s[0])
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i]) //相应Unicode码点所表示的字面量
	}
	fmt.Println()

	//strings
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Index("china", "a"))
	fmt.Println(strings.Index("china", "b"))
	fmt.Println(strings.TrimSpace(" I don ’ t need all this space ")) //删除开头和末尾的空白

	//函数和方法的名称可包含UTF-8 支持的任何字符, so as PHP.
	函数()

	//字符串是不可变的
	//s[2] = "h"
	var new_s []string
	for i := 0; i < len(s); i++ {
		if i == 2 {
			new_s = append(new_s, "L")
		} else {
			new_s = append(new_s, string(s[i]))
		}
	}
	fmt.Println(myToString(new_s))
	fmt.Println(strings.Replace(s, "l", "L", 1))
}

func 函数() {
	fmt.Println("函数")
}
func myToString(s []string) string {
	var new_s string
	for i := 0; i < len(s); i++ {
		new_s += string(s[i])
	}
	return new_s
}
