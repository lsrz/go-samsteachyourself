package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {

	//单引号 single quotes
	//	单引号只能有一个字符，如果输出会返回这个字符的ascii码 ，如果想输出为字符需要用string()函数转换一下
	//	单引号，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。
	sa := 'a'
	fmt.Println(sa)
	fmt.Println(string(sa))
	fmt.Printf("Size: %d\nType: %s\nUnicode CodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(sa), reflect.TypeOf(sa), sa, sa)

	var sb byte = 'a'
	fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(sb), reflect.TypeOf(sb), sb)

	//双引号 double quotes
	//	Go中字符串是一个不可变的值类型，内部用指针指向UTF-8字节数组
	//	双引号，才是字符串，实际上是字符数组。可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。
	da := "abcdefg"
	fmt.Println(da[0])   //97
	fmt.Println(da[0:2]) //ab

	db := "hello," + //拼接符需要放在上一行的末尾，这是因为编译器会进行行尾自动补全分号的缘故。
		"world"
	fmt.Println(db)

	dc := []string{"hello", "world"}
	dd := strings.Join(dc, ",")
	fmt.Println(dd)

	df := "tit\nfor\ttat"
	fmt.Printf("double quotes is: %s\n", df)

	//反引号 back quotes
	//	表示字符串字面量，但不支持任何转义序列。字面量 raw literal string 的意思是，你定义时写的啥样，它就啥样，你有换行，它就换行。你写转义字符，它也就展示转义字符。
	ba := `tit\nfor\ttat`
	fmt.Printf("back quotes is: %s\n", ba)
}
