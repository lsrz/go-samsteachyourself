package main

import (
	"fmt"
	"strings"
)

/*

常用的测试
	1 单元测试：针对一小部分代码，并独立地对它们进行测试，如函数测试。
	2 集成测试：测试的是应用程序各部分协同工作的情况，集成测试还检查诸如网络调用和数据库连接等方面，以确保整个系统按期望的那样工作。
	3 功能测试：从最终用户的角度核实软件按期望的那样工作，它们评估从外部看到的程序的运行情况，而不关心软件内部的工作原理。

测试驱动开发(TDD)[是用测试驱动开发，不是测试驱动的开发 WTF]
	从测试的角度考虑新功能，先编写测试来描述代码片段的功能，再着手编写代码。

testing包
	约定一：同目录_testing结尾
	约定二：测试函数、方法以Test开头
		func TestFunc(t *testing.T) {...

	go test  //命令测试整个package,可使用doctest测试
	go test -v
	go test example01_test.go
		got want 模式

运行表格驱动测试
	结构体切片表格
	参考example02_test.go

运行基准测试
	参考example03_test.go
		func BenchmarkFunc(b *testing.B) {
			for n := 0; n < b.N; n++ {...

	go test -bench=.
	go test -bench=. example03_test.go

		goos: windows
		goarch: amd64
		pkg: doctest
		cpu: Intel(R) Core(TM) i3-8100 CPU @ 3.60GHz
		BenchmarkStringFromAssignment-4           188648              6403 ns/op
		BenchmarkStringFromAppendJoin-4           546027              2242 ns/op
		BenchmarkStringFromBuffer-4              1286613               939.6 ns/op
		PASS
		ok      doctest 4.943s

		含义：
		函数后面的-4了吗？这个表示运行时对应的GOMAXPROCS的值。
		接着的188648表示运行for循环的次数也就是调用被测试代码的次数。
		最后的6403ns/op表示每次需要话费6403纳秒。(执行一次操作花费的时间)

		1s(秒)=10^3ms(毫秒)=10^6μs(微秒)=10^9ns(纳秒)=10^12ps(皮秒)=10^15fs(飞秒)=10^18as(阿秒)=10^21zm(仄秒)=10^24ym(幺秒)秒
覆盖率
	go test -cover
	go test -cover example01_test.go

	实现100% 的测试覆盖率是一个值得为之努力的目标，但对大型项目而言，这几乎是不可能的。
	达到8 0% 左右的测试覆盖率就可以了， 具体多少取决于项目的复杂度。

*/
func main() {
	fmt.Println(strings.ToUpper("hello world"))
}
