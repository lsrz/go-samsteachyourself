package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	//使用了标准库strings 将其转换为io.Reader
	postData := strings.NewReader(`{"some":"json"}`)
	//指定超时时间
	client := &http.Client{
		Timeout: 1 * time.Microsecond,
	}
	response, err := client.Post("https://httpbin.org/post", "application/json", postData)
	//http.Post
	//response, err := http.Post("https://httpbin.org/post", "application/json", postData)
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
