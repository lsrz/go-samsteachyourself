package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"github.com/lsrz/go/temperature"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(strings.ToUpper("hello"))

	//go get github.com/golang/example/stringutil
	// 命令go get 很聪明，它会下载依赖的第三方包
	// 这个包被安装到环境变量GO PATH 指定的路径中
	// %GOPATH%/src/..
	s := "ti esrever dna ti pilf nwod gniht ym tup I"
	fmt.Println(stringutil.Reverse(s))

	//更新特定的包
	// go get -u github.com/spf13/hugo

	//更新文件系统中所有的包
	// go get -u all

	//vendor
	// go1.5版引入了文件夹vendor，这能够让您将第三方模块添加到项目目录下的文件夹vendor中，并将所有包文件都移到这个文件夹中。
	// go1.6以后编译go代码会优先从vendor目录先寻找依赖包,好处：锁定版本

	//创建包
	// 包中方法和文件名无关
	// go get github.com/lsrz/go/temperature
	temperature.Show()
	//大小写问题：
	// 以大写字母打头的标识符会被导出，这意味着导入包后就可使用它们；
	// 以小写字母打头的标识符不会被导出，这意味着即便导入包也无法使用它们。
	// 简而言之，以大写字母打头的标识符是公有的，而以小写字母打头的标识符是私有的。
	//temperature.smallShow()
	temper := temperature.Measure()
	fmt.Println(temper)

	/*
		go get为什么不下载源代码到src下?
	
		使用命令查看go mod的功能是否开启： go env
		默认情况下可以看到GO111MODULE=""，即GO111MODULE=auto
		要想go get下载源代码到src下，设置关闭go mod功能既可以了：
		windows下： go env -w GO111MODULE=off
		Linux下： export GO111MODULE=off
	*/
}
