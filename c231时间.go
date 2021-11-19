package main

import (
	"fmt"
	"log"
	"time"
)

//ticker [ˈtɪkə(r)] 滴答响的东西;时钟;振动子

func main() {
	//2021-11-19 14:09:20.8516591 +0800 CST m=+0.002869901
	//网络时间协议（Network Time Protocol, NTP）
	//世界标准时间（Coordinated Universal Time, UTC）。UTC是时间标准而非时区
	//CST可视为美国、澳大利亚、古巴或中国的标准时间
	//单调时钟（monotonic clock）表示从过去某个任意的固定点开始经过的绝对时间
	//If the time has a monotonic clock reading, the returned string includes a final field "m=±<value>", where value is the monotonic clock reading formatted as a decimal number of seconds.

	//time.Now()
	fmt.Println(time.Now())

	//time.Sleep()
	time.Sleep(3 * time.Second)
	fmt.Printf("i'm awake at:\n%v\n", time.Now())

	//time.Parse()
	s := "2006-01-02T15:04:05+07:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	//Time结构体
	fmt.Printf("The hour is %v\n", t.Hour())
	fmt.Printf("The minute is %v\n", t.Minute())
	fmt.Printf("The second is %v\n", t.Second())
	fmt.Printf("The day is %v\n", t.Day())
	fmt.Printf("The month is %v\n", t.Month())
	fmt.Printf("The day of the week is %v\n", t.Weekday())
	fmt.Printf("UNIX time is %v\n", t.Unix())

	//时间加减
	//t.Add
	at := t.Add(2 * time.Second)
	fmt.Println(at)
	//t.Sub
	st := t.Sub(at)
	fmt.Println(st)

	//比较两个不同的Time结构体
	s1 := "2017-09-03T18:00:00+00:00"
	s2 := "2017-09-04T18:00:00+00:00"
	today, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}
	tomorrow, err := time.Parse(time.RFC3339, s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
	fmt.Println(today.Equal(tomorrow))

	//time.After()
	//特定的时间过后执行某项操作(向通道发送一条消息)
	//for{
	for i := 0; i < 3; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("2 seconds time's up !")
			//return //the return is for main func
			break
		}
	}
	//time.Tick()
	//特定的时间过后执行某项操作
	c := time.Tick(2 * time.Second)
	for t := range c {
		fmt.Printf("The time is now %v\n", t)
	}
}
