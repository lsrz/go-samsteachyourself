package main

import (
	//标准库中的net/http包提供了多种创建HTTP服务器的方法
	"net/http"
)

//接受一个http.ResponseWriter和一个指向请求的指针
func DefaultIndex(w http.ResponseWriter, r *http.Request) {

	//Write将－个字节切片作为参数
	w.Write([]byte ("Hello World\n"))
}

func main() {

	//HandleFunc创建了路由／，这个方法接受一个模式和一个函数
	http.HandleFunc("/", DefaultIndex)

	//ListenAndServe 来启动一个服务器
	http.ListenAndServe(":8000", nil)
}
