package main

import (
	"fmt"
	"os"
)

/*
https://pkg.go.dev/fmt
fmt is short for format.
fmt.Printf("some desc %v\n", v)

Printf 格式化输出
https://blog.csdn.net/qq_34777600/article/details/81266453

通用占位符：
	v     值的默认格式。
	%+v   添加字段名(如结构体)
	%#v　 相应值的Go语法表示
	%T    相应值的类型的Go语法表示
	%%    字面上的百分号，并非值的占位符
布尔值：
	%t   true 或 false
整数值：
	%b     二进制表示
	%c     相应Unicode码点所表示的字符
	%d     十进制表示
	%o     八进制表示
	%q     单引号围绕的字符字面值，由Go语法安全地转义
	%x     十六进制表示，字母形式为小写 a-f
	%X     十六进制表示，字母形式为大写 A-F
	%U     Unicode格式：U+1234，等同于 "U+%04X"
浮点数及复数：
	%b     无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat中的 'b' 转换格式一致。例如 -123456p-78
	%e     科学计数法，例如 -1234.456e+78
	%E     科学计数法，例如 -1234.456E+78
	%f     有小数点而无指数，例如 123.456
	%g     根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	%G     根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
字符串和bytes的slice表示：
	%s     字符串或切片的无解译字节
	%q     双引号围绕的字符串，由Go语法安全地转义
	%x     十六进制，小写字母，每字节两个字符
	%X     十六进制，大写字母，每字节两个字符
指针：
	%p     十六进制表示，前缀 0x
*/

/*
对于％ｖ来说默认的格式是：
	bool:                    %t
	int, int8 etc.:          %d
	uint, uint8 etc.:        %d, %x if printed with %#v
	float32, complex64, etc: %g
	string:                  %s
	chan:                    %p
	pointer:                 %p
*/

func main() {
	err := fmt.Errorf("%v\n", "Errorf")
	n, errs := fmt.Printf("use Printf output %s", err)
	fmt.Println("PrintXX returns:", n, errs)

	sth := fmt.Sprintf("%v\n", "Sprintf")
	fmt.Fprintf(os.Stdout, "use Fprintf output %s\n", sth)
}

//Errorf

//PrintXX返回n,error
//Print,Pringln,Printf

//SprintXX返回string
//Sprint,Sprintln,Sprintf

//FprintXX把内容写到响应流中
//Fprint,Fprintln,Fprintf

//Errorf返回string
