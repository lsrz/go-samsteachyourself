package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	/*
		# For OSX and Linux
		$ DEBUG=1 go run xx.go
		# For Windows
		$ set DEBUG=1
		$ go run xx.go
	*/
	debug := os.Getenv("DEBUG")
	fmt.Printf("DEBUG: %s\n", debug)

	//创建http.Client客户端
	//！！！Timeout不起作用！！！
	client := &http.Client{
		Timeout: 1 * time.Microsecond, //处理超时
	}
	/*
		//通过创建一个传输（ transport ）并将其传递给客户端
		tr := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   50 * time.Millisecond,
				KeepAlive: 50 * time.Millisecond,
			}).DialContext,
			TLSHandshakeTimeout:   50 * time.Millisecond,
			IdleConnTimeout:       50 * time.Millisecond,
			ResponseHeaderTimeout: 50 * time.Millisecond,
			ExpectContinueTimeout: 50 * time.Millisecond,
		}
		client := &http.Client{
			Transport: tr,
		}
	*/

	//http.NewRequest
	request, err :=http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	//add some headers
	//request.Header.Add("Accept", "application/json")

	if debug == "1" {
		//httputil.DumpRequestOut
		/*
			GET / HTTP/1.1
			Host: 127.0.0.1
			User-Agent: Go-http-client/1.1
			Accept-Encoding: gzip
		*/
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("debugRequest: \n%s\n", debugRequest)
	}

	//client.Do
	response, err := client.Do(request)
	defer response.Body.Close()

	if debug == "1" {
		//httputil.DumpResponse
		/*
			HTTP/1.1 200 OK
			Content-Length: 26
			Connection: Upgrade
			Content-Type: text/html; charset=UTF-8
			Date: Mon, 15 Nov 2021 08:14:15 GMT
			Server: Apache/2.4.48 (Win64) OpenSSL/1.1.1k PHP/7.4.24
			Upgrade: h2,h2c
			X-Powered-By: PHP/7.4.24

			this is a test index page
		*/
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("DumpResponse: \n%s\n", debugResponse)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ioutil.ReadAll:\n%s\n", body)
}
