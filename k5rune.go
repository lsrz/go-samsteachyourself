package main

import (
	"fmt"
	"unicode/utf8"
)

/*
rune [ruːn]符文、字母

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
// int32的别名，几乎在所有方面等同于int32，它用来区分字符值和整数值
type rune = int32

golang默认编码是utf-8
golang中string底层是通过byte数组实现的。
中文字符在unicode下占2个字节，
中文字符在utf-8编码下占3个字节。
*/

func main() {

	var str = "hello 世界"

	//golang中string底层是通过byte数组实现的，直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))

	fmt.Println("rune:", []rune(str))
}
