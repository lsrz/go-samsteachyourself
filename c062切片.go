package main

import (
	"fmt"
	"reflect"
)

//slice[slaɪs]切片
//除非确定必须使用数组，否则请使用切片
func main() {

	//var cheeses1 []string

	var cheeses1 = make([]string, 3)
	cheeses1[0] = "a"
	cheeses1[1] = "b"
	fmt.Println(cheeses1)

	//切片类似于数组，但不同于数组的是，您可在切片中添加和删除元素。
	//append会在必要时调整切片的长度，但它对程序员隐藏了这种复杂性。

	//添加一个元素
	//append从切片的末尾开始写入，注意此时为a，b，空，d。长度4
	cheeses1 = append(cheeses1, "d")
	fmt.Println(cheeses1, len(cheeses1))

	//添加多个元素
	cheeses1 = append(cheeses1, "e", "f", "g")
	fmt.Println(cheeses1)

	//删除第n个元素(n从0开始)
	//sth = append (sth[:n] , sth[n+1:]...) 没有专门用于从切片中删除元素的函数
	var cheeses2 = make([]string, 3)
	cheeses2[0] = "Mariolles"
	cheeses2[1] = "Epoisses de Bourgogne"
	cheeses2[2] = "Camembert"
	fmt.Printf("切片长度: %v\n", len(cheeses2))
	fmt.Println(cheeses2)
	n := 0
	cheeses2 = append(cheeses2[:n], cheeses2[n+1:]...)
	fmt.Printf("切片长度: %v\n", len(cheeses2))
	fmt.Println(cheeses2)

	n = 0
	cheeses2 = slice_delete(cheeses2, n)
	fmt.Printf("切片长度: %v\n", len(cheeses2))
	fmt.Println(cheeses2)

	//复制
	var cheeses3 = make([]string, 3)
	copy(cheeses3, cheeses2)
	fmt.Println("Copy:", cheeses3, "Len:", len(cheeses3))

	//修改
	cheeses3[0] = "new Camembert"
	fmt.Println("Edit:", cheeses3)
	fmt.Println("Type:", reflect.TypeOf(cheeses3))

	//创建长度和容量都为3的slice，和数据的区别是中括号无数字
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	//这叫字节，区分字面量
	byte1 := []byte{'a', 'a', 'a'}
	fmt.Println(byte1)
	//!!!括号是类型转换!!!
	byte2 := []byte("aaa")
	fmt.Println(byte2)
	for _, v := range byte2 {
		fmt.Printf("%s,%s\n", v, string(v))
	}
	byte3 := []byte(`{"hello":"world"}`)
	for _, v := range byte3 {
		fmt.Printf("%s,%s\n", v, string(v))
	}
}

func slice_delete(m []string, n int) []string {
	return append(m[:n], m[n+1:]...)
}
