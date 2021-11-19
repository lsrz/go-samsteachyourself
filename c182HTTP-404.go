package main

import (
	//标准库中的net/http包提供了多种创建HTTP服务器的方法
	"net/http"
)

//接受一个http.ResponseWriter和一个指向请求的指针
func DefaultIndexPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	//
	w.Header().Set("X My-Header Before", "I am setting a header before write")

	//Write将－个字节切片作为参数
	//声明一个字节切片并将字符串值转换为字节
	w.Write([]byte ("Hello World\n"))

	//发送响应是最后一步，鉴于响应己经发送，这行代码不会有任何作用，但能够通过编译
	w.Header().Set("X My-Header After", "I am setting a header after write")
}

func main() {

	//HandleFunc创建了路由／，这个方法接受一个模式和一个函数
	//路由器默认将没有指定处理程序的请求定向到／，localhost:8000/user
	//默认路由不支持变量，/product/:id
	http.HandleFunc("/", DefaultIndexPage)

	//ListenAndServe 来启动一个服务器
	http.ListenAndServe(":8000", nil)
}
