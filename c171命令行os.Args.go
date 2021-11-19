package main

import (
	"fmt"
	"os"
)

/*
标准输入 0
标准输出 1
标准错误 2
*/
func main() {
	//访问命令行参数
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
	// go run c171命令行os.Args.go 1 2 3
	//0 C:\..Temp\go-build..\c171命令行.exe
	//1 1
	//2 2
	//3 3
}
