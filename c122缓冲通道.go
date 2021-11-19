package main

import "fmt"

/*
	缓冲意味着可将数据存储在通道中，等接收者准备就绪再交给它。
*/

/*
var message = make(chan string, 2)


func Worker1() {
	message <- "hello"
}

func Worker2() {
	message <- "world"
}
*/
func receiver(c chan string) {
	for msg := range c {
		fmt.Println(string(msg))
	}
}
func main() {

	/*
		go Worker1()
		go Worker2()

		//只能接到1条？(它将阻止函数返回，直到收到一条消息为止)
		//msg := <-message
		//仍然要等待
		time.Sleep(time.Second * 2)
		close(message)

		for msg := range message {
			fmt.Println(string(msg))
		}
		//above all 一定对接收方有什么误解
	*/

	var message = make(chan string, 2)
	message <- "hello"
	message <- "world"
	close(message)

	receiver(message)
}
