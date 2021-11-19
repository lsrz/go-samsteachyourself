package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//JavaScript对象表示法(JavaScript Object Notation, JSON ）
//己成为互联网上存储和交换数据的事实标准，从很大程度上说它己取代可扩展的标记语言（ Extensi ble Markup Language, XML) 。
//要编码或解码JSON，必须创建结构体
/*
{
	"name"		:	"George",
	"age"		:	40,
	"children"	:	["Bea","Fin"]
}
*/

func main() {
	hobbies := []string{"Cycling", "Cheese", "Techno"}

	type Persons struct {
		Name    string   `json:"name,omitempty"`
		Age     int      `json:"age,omitempty"`
		Hobbies []string `json:"hobbies,omitempty"` //omitempty = omit[əˈmɪt]忽略 + empty
	}
	type Switch struct {
		On bool `json:"on"`
	}

	//json字段key是小写的，使用go标签替换
	//omitempty,忽略零值（"",0,nil）
	p1 := []Persons{
		{
			Name:    "George",
			Age:     40,
			Hobbies: hobbies,
		},
		{
			Name:    "",
			Age:     40,
			Hobbies: hobbies,
		},
		{
			Name:    "George",
			Age:     0,
			Hobbies: hobbies,
		},
		{
			Name:    "George",
			Age:     40,
			Hobbies: nil,
		},
	}

	fmt.Printf("%+v\n", p1)

	//结构体编码为JSON(字节切片),[ˈmɑːʃl]
	jsonByteData, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	//输出JSON
	fmt.Println(string(jsonByteData))

	//JSON解码为结构体
	var jsonStringData = `{"name":"George", "age":40, "hobbies":["Cycling", "Cheese", "Techno"]}`
	jsonByteData = []byte(jsonStringData)
	p2 := Persons{}
	err = json.Unmarshal(jsonByteData, &p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p2)
	fmt.Printf("%v\n", p2.Name)

	//映射注意类型
	var jsonString = `{"on":"true"}`
	jsonByte := []byte(jsonString)
	s1 := Switch{}
	err = json.Unmarshal(jsonByte, &s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s1)

}
