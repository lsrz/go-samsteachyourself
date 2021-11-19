package main

/*
Go目标：兼具Python等动态语言开发速度和C或C++等编译型语言的性能和安全性

设计GO语言是因为C++给他们带来了挫败感，厌烦了等待编译完成。

GO是编译型语言

go version
go evn

GOPATH                #账户的环境变量
echo %GOPATH%

mkdir %GOPATH%\bin   #可执行文件
mkdir %GOPATH%\pkg   #归档文件
mkdir %GOPATH%\src   #源码文件


GOROOT  表示SDK的存放路径
GOPATH  表示代码的存放路径
GOLAND  表示开发环境的存放路径

----------------------------------------------------------
GoLand2020.2 找不到go1.17 SDK,升级到GoLand2020.3或降级到go1.16
----------------------------------------------------------

go build  hello.go #测试编译
go run    hello.go #编译并运行，开发GO时，没有必要将编译及执行步骤分开
go clean  hello.go #清除编译文件
go install         #编译并安装，将编译的中间文件放在GOPATH的pkg目录，将编译结果放在GOPATH的bin目录

hello.exe

gopher  ˈɡəʊfə(r)   囊鼠
*/

/*

深入理解 go build 和 go install

1.作用
	go build：	用于测试编译包，在项目目录下生成可执行文件（有main包）。
	go install：主要用来生成库和工具。
				一是编译包文件（无main包），将编译后的包文件放到 pkg 目录下（$GOPATH/pkg）。
				二是编译生成可执行文件（有main包），将可执行文件放到 bin 目录（$GOPATH/bin）。
2.相同点
	都能生成可执行文件

3.不同点
	go build 不能生成包文件
	go install 可以生成包文件

	go build 生成可执行文件在当前目录下
	go install 生成可执行文件在bin目录下（$GOPATH/bin）

*/
