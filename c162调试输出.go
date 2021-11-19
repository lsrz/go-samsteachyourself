package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//打印输出
	//添加打印数据能快速发现Bug。然而，使用这种方法将导致代码中充斥调试代码。
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input a name:")
	s, _ := reader.ReadString('\n')

	//s = strings.TrimSpace(s)
	s = strings.Replace(s, "\n", "", -1)

	fmt.Printf("[debug] input text: %v\n", s)

	if s == "GO" {
		fmt.Println("you are right")
	} else {
		fmt.Println("you are wrong")
	}
}
