package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	//(?i) 不区分大小写
	patten := "(?i)chocolate"
	str := "Chocolate is my favorite !"
	//MatchString
	matched, err := regexp.MatchString(patten, str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(matched)

	/*
		字符		含义
		. 		与除换行符之前的其他任何字符都匹配
		* 		与苓个或多个指定的字符匹配
		^ 		表示行首
		$ 		表示行尾
		+ 		匹配一次或多次
		? 		匹配零或一次
		[] 		与方摇号内指定的任何字符都匹配
		{n} 	匹配n次
		{n,} 	匹配n次或更多次
		{m,n} 	最少匹配m次，最多匹配n次
	*/
	/*
	   [:alnum:] 匹配任何字母
	   [:alpha:] 匹配任何字母和数字
	   [:digit:] 匹配任何数字
	   [:upper:] 匹配任何大写字母
	   [:lower:] 匹配任何小写字母
	   [:punct:] 匹配任何标点符号
	   [:space:] 匹配空格符
	   [:xdigit:] 匹配任何16进制数字
	   [:blank:] Blank characters: space and tab.
	   [:cntrl:] Control characters. In ASCII, these characters have octal codes 000 through 037, and 177 (DEL). In other character sets, these are the equivalent characters, if any.
	   [:graph:] Graphical characters: [:alnum:] and [:punct:].
	   [:print:] Printable characters: [:alnum:], [:punct:], and space.
	*/

	//要将正则表达式赋给变量，必须先对其进行分析。
	//	Compile ：未能通过编译时返回错误。
	//	MustCompile ：无法编译时引发panic
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdf£33£3"))
	fmt.Println(re.MatchString("roger"))
	fmt.Println(re.MatchString("iamthebestuseofthisappevaaaar"))

	//替换
	usernames := [4]string{
		"slimshady99",
		"!asdf£33£3",
		"roger",
		"iamthebestuseofthisappevaaaar",
	}
	an := regexp.MustCompile("[^[:alnum:]]")

	for _, username := range usernames {
		if len(username) > 12 {
			username = username[:12]
			fmt.Printf("trimmed username to %v\n", username)
		}
		//ReplaceAllString
		if !re.MatchString(username) {
			username = an.ReplaceAllString(username, "x")
			fmt.Printf("rewrote username to %v\n", username)
		}
	}
}
