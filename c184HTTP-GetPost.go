package main

import (
	//标准库中的net/http包提供了多种创建HTTP服务器的方法
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DealGetPost(w http.ResponseWriter, r *http.Request) {

	rMethod := r.Method

	switch rMethod {

	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf(" %s: %s\n ", k, v)
		}
		w.Write([]byte("Received a GET request\n"))
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" %s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte (http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {

	http.HandleFunc("/", DealGetPost)

	http.ListenAndServe(":8000", nil)
}
