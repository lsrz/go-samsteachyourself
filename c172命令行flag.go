package main

import (
	"flag"
	"fmt"
	"os"
)

//创建命令
func main() {

	//自定义帮助文本
	flag.Usage = func() {
		usageText := `Usage c172命令行flag [OPTION]
An example of customizing usage output

-s, --s example string argument, default: String help text
-i, --i example integer argument, default: Int help text
-b, --b example boolean argument, default: Bool help text`
		fmt.Fprintf(os.Stderr, "%s\n", usageText)
	}

	//flag分析命令行标志
	// 比os.Args能多起名字
	a := flag.String("s", "default value", "String s")
	b := flag.Int("i", 0, "Int i")
	c := flag.Bool("b", false, "Bool b")

	flag.Parse() //调用flag.Parse，让程序能够传递声明的参数。

	fmt.Println(*a, *b, *c)
	/*
		go run c172命令行flag.go
		default value

		go run c172命令行flag.go -s a
		a

		go run c172命令行flag.go -d a
		flag provided but not defined: -d
		Usage of C:\...\c172命令行flag.exe:
		-s string
		introductions (default "default value")
		exit status 2

		go run c172命令行flag.go -s a -i 2
		a 2

		go run c172命令行flag.go --s a --i 2
		虽然使用一个连字符和两个连字符的选项通常是不同的选项，但Go设计者将它们视为相同的选项。
	*/
}
