package main

import (
	"log"
	"os"
)

func main() {
	//1. 日志
	//log日志
	log.Printf("This is a log message")

	//致命错误（with exit code 1）
	//var errFatal = errors.New("We only just started and we are crashing")
	//log.Fatal(errFatal)

	//将日志写入文件
	f, err := os.OpenFile("c161调试.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.SetOutput(f)

	for i := 1; i <= 5; i++ {
		log.Printf("Log iteration %d", i)
	}

	//通常，最好使用操作系统将日志重定向到文件，而不要使用Go语言代码。因为这种方法更灵活，能够让其他工具在必要时使用日志。
	// go run c161调试.go > c161调试.log 2>&1
}
