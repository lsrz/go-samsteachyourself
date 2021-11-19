package main

import "fmt"

func main() {
	//Delve提供了更精致的程序调试环境
	//go get github.com/derekparker/delve/cmd/dlv
	//go install github.com/go-delve/delve/cmd/dlv@v1.7.2

	//dlv debug c164调试delve.go
	s := "Hello World"
	t := "Goodbye, Cruel World "
	fmt.Printf("s is %v, t is %v\n", s, t)
}
