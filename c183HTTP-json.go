package main

import (
	//标准库中的net/http包提供了多种创建HTTP服务器的方法
	"net/http"
)

func DefaultIndexJSON(w http.ResponseWriter, r *http.Request) {

	//获取发送Header
	rheader := r.Header.Get("Accept")
	//rheader := r.Header.Get("Accept-Encoding")

	//设置返回Header
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//返回JSON格式
	w.Write([]byte (`{"hello":"world";"header": "` + rheader + `";}`))
}

func main() {

	http.HandleFunc("/", DefaultIndexJSON)

	http.ListenAndServe(":8000", nil)
}
