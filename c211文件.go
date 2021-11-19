package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

//ioutil - util跑龙套
//ioutil 就是一个OS模块包装器，使用它需要编写的代码更少，且无须执行清理操作。

func main() {

	//ReadFile
	fileBytes, err1 := ioutil.ReadFile("c161调试.log")
	if err1 != nil {
		log.Fatal(err1)
	}

	fileString := string(fileBytes)
	fmt.Println(fileString)

	//WriteFile
	var b = make([]byte, 0)
	err2 := ioutil.WriteFile("example02.txt", b, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

	//ReadDir
	file, err3 := ioutil.ReadDir(".")
	if err3 != nil {
		log.Fatal(err3)
	}
	for k, v := range file {
		fmt.Println(k, v.Name())
	}

	//Copy（不同的操作性系统差别很大，这导致创建通用的文件复制方法是很难的。鉴于此，没有用于复制文件的便利方法，要复制文件，应使用OS包）
	from, err4 := os.Open("example02.txt")
	if err4 != nil {
		log.Fatal(err4)
	}
	defer from.Close()
	to, err5 := os.OpenFile("example02_copy.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err5 != nil {
		log.Fatal(err5)
	}
	defer to.Close()

	_, err6 := io.Copy(to, from)
	if err6 != nil {
		log.Fatal(err6)
	}

	//Remove
	FileName := "test.txt"
	info, err := os.Stat(FileName)
	if err != nil {
		//文件不存在
		fmt.Printf("File Does Not Exists : \"%s\"\n", FileName)
	} else {
		fmt.Println(info)
		err7 := os.Remove("test.txt")
		if err7 != nil {
			log.Fatal(err7)
		}
	}

	//json config
	type Config struct {
		Version float64 `json:"version"`
		Name    string  `json:"name"`
	}
	FileBytes, err8 := ioutil.ReadFile("config.json")
	if err8 != nil {
		log.Fatal(err8)
	}
	c := Config{}
	err9 := json.Unmarshal(FileBytes, &c)
	if err9 != nil {
		log.Fatal(err9)
	}
	fmt.Printf("%+v\n", c)

	//toml config
	type Config2 struct {
		Age        int
		Cats       []string
		Pi         float64
		Perfection []int
		DOB        time.Time // requires `import time`
	}
	t := Config2{}
	_, err10 := toml.DecodeFile("config.toml", &t)
	if err10 != nil {
		log.Fatal(err10)
	}
	fmt.Printf("%+v\n", t)
}
