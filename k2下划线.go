package main

/*
在 Golang 里， _ （下划线）是个特殊的标识符。

用在 import

	在导包的时候，常见这个用法，尤其是项目中使用到 mysql 或者使用 pprof 做性能分析时，比如
	import _ "net/http/pprof"
	import _ "github.com/go-sql-driver/mysql"
	这种用法，会调用包中的init()函数，让导入的包做初始化，但是却不使用包中其他功能。

用在返回值

	该用法也是一个常见用法。Golang 中的函数返回值一般是多个，err 通常在返回值最后一个值。
	但是，有时候函数返回值中的某个值我们不关心，如何接收了这个值但不使用，代码编译会报错，因此需要将其忽略掉。比如
	for _, val := range Slice {}
	_, err := func()

用在变量

	type I interface {
		Sing()
	}

	type T struct {
	}

	func (t T) Sing() {
	}

	// 编译通过
	var _ I = T{}
	// 编译通过
	var _ I = &T{}

	在这里下划线用来判断结构体是否实现了接口，如果没有实现，在编译的时候就能暴露出问题，如果没有这个判断，后代码中使用结构体没有实现的接口方法，在编译器是不会报错的。
*/


