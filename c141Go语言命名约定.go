package main

import "fmt"

//标准约定

//gofmt格式
//gofmt test.go
//gofmt -w test.go

//大小写、驼峰
//var FileName:=""
//var fileName:="" //不被导出

//golint 提供了有关风格方面的提示
// golint test.go
/*
	git clone https://github.com/golang/lint.git
	cd lint/golint
	go install
	D:\GoLang\GoProjects\bin\golint.exe
	https://blog.csdn.net/qq_34326321/article/details/105122219
*/

//godoc
// go install golang.org/x/tools/cmd/godoc@latest
// go doc命令可以打印附于Go语言程序实体上的文档
// 		go doc DocTest
// 		go doc DocTest SomeMethod
// godoc是一个很强大的工具，同样用于展示指定代码包的文档
// 		godoc -http=localhost:6060

//工作流程自动化
// -l list files whose formatting differs from gofmt's
//gofmt -l .

//在给接口命名方面，Go语言惯用的做法是什么？
//在接口命名方面，惯用的做法是使用一个动词，并在末尾加上er，如Parser和Authorizer，这样的名称描述了接口是做什么的。

func main() {
	fmt.Println("ftm test")
	fmt.Println("ftm test new line")
}

/*

Go中的命名规范整理
	https://www.cnblogs.com/zhangyafei/p/12466162.html
	https://blog.csdn.net/flysnow_org/article/details/103520993

一、package命名规范
    包名用小写,使用短命名,尽量和标准库不要冲突。
	包名统一使用单数形式。

二、文件命名规范
    文件命名一律采用小写，不用驼峰式，尽量见名思义，看见文件名就可以知道这个文件下的大概内容。
    其中测试文件以_test.go结尾，除测试文件外，命名不出现_。

	如：stringutil.go， stringutil_test.go

三、变量命名规范
	变量命名一般采用驼峰式，当遇到特有名词（缩写或简称，如DNS）的时候，特有名词根据是否私有全部大写或小写。

	如：apiClient， URLString

四、常量命名规范
    同变量规则，力求语义表达完整清楚，不要嫌名字长。
    如果模块复杂，为避免混淆，可按功能统一定义在package下的一个文件中。

	如： const todayNews = "Hello"
		// 如果超过了一个常量应该用括号的方法来组织
		const (
		  systemName = "What"
		  sysVal = "dasdsada"
		)

五、函数/方法命名规范
	由于Golang的特殊性（用大小写来控制函数的可见性），除特殊的性能测试与单元测试函数之外, 都应该遵循如下原则
	1. 采用驼峰式。将功能及必要的参数体现在名字中，不要嫌长，如updateById，getUserInfo.
	2. 如果包外不需要访问请用小写开头的函数
	3. 如果需要暴露出去给包外访问需要使用大写开头的函数名称

	// 对象暴露的方法
	func (*fileDao) AddFile(file *model.File) bool { ...

	// 不需要给包外访问的函数
	func removeCommaAndQuote(content string) string { ...

六、结构体命名规范
    结构体名应该是名词或名词短语，如WikiPage、Account，避免使用Manager、Processor、Data、Info等类名，而且类名不应当是动词。
	结构体中属性名大写，如果属性名小写则在数据解析（如json解析,或将结构体作为请求或访问参数）时无法解析

	type Webhook struct {
		ID           int64
		RepoID       int64
		OrgID        int64
		URL          string
	}

七、注释命名规范
	每个包都应该有一个包注释，位于 package 之前。
	如果同一个包有多个文件，只需要在一个文件中编写即可；
	如果你想在每个文件中的头部加上注释，需要在版权注释和 Package前面加一个空行，否则版权注释会作为Package的注释。

	// Copyright 2009 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.
	package net

	每个以大写字母开头（即可以导出）的方法应该有注释，且以该函数名开头。

	// Get 会响应对应路由转发过来的 get 请求
	func (c *Controller) Get() {...

	大写字母开头的方法以为着是可供调用的公共方法，如果你的方法想只在本包内掉用，请以小写字母开发。

	func (c *Controller) curl() {...

	注释应该用一个完整的句子，
	注释的第一个单词应该是要注释的指示符，以便在 godoc 中容易查找。
	注释应该以一个句点 . 结束。

八、接口命名规范
    单个函数的接口名以 er 为后缀
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	两个函数的接口名综合两个函数名
	type WriteFlusher interface {
		Write([]byte) (int, error)
		Flush() error
	}

	三个以上函数的接口名类似于结构体名
	type Car interface {
		Start()
		Stop()
		Drive()
	}

九、receiver命名规范
    golang 中存在receiver 的概念 Receiver 的名称应该缩写，应该尽量保持一致， 避免this, super，等其他语言的一些语义关键字如下

	type A struct{}
	func (a *A) methodA() {
	}
	func (a *A) methodB() {
	  a.methodA()
	}

十、import 规范
	如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，建议采用如下方式进行组织你的包：
	有顺序的引入包，不同的类型采用空格分离，第一种实标准库，第二是项目包，第三是第三方包。

	import (
		"encoding/json"
		"strings"
		"fmt"

		"myproject/models"
		"myproject/controller"
		"myproject/utils"

		"github.com/astaxie/beego"
		"github.com/go-sql-driver/mysql"
	)

	在项目中不要使用相对路径引入包：
	import "…/net"
*/
