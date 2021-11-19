package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//指定超时时间
	//因为每条连接都会占用一些内存，还会使用底层操作系统中的一个套接字(即Socket=IP地址:端口号)。
	//如果连接的速度很慢，则程序很快就会出现内存泄露，或耗尽底层操作系统的资源。
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	response, err := client.Get("http://127.0.0.1")
	//http.Get
	//response, err := http.Get("http://127.0.0.1")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
