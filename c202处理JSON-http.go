package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	type Result struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	}

	client := http.Client{
		Timeout: 3 * time.Second,
	}
	//获取数据
	httpResult, err := client.Get("https://api.github.com/users/shapeshed")
	if err != nil {
		log.Fatal(err)
	}
	defer httpResult.Body.Close()

	//方法一：得到字符串
	//body, err := ioutil.ReadAll(httpResult.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//httpByteResult := []byte(body)
	//
	//r1 := Result{}
	//err = json.Unmarshal(httpByteResult, &r1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%+v", r1)

	//方法二：
	var r2 Result
	err = json.NewDecoder(httpResult.Body).Decode(&r2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", r2)
}
