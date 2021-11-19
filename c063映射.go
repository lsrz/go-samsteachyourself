package main

import "fmt"

//其实就是key value字典
//value类型必须相同
//无长度,第二参数只是容量提示，而非硬性规定
//有删除函数

func main() {
	var players = make(map[string]int)

	players["lily"] = 24
	players["lucy"] = 25
	players["hanm"] = 26

	fmt.Println(players)

	//删除映射
	delete(players, "lily")
	fmt.Println(players)

}
