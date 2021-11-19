package main

import (
	"errors"
	"fmt"
)

//type sth interface
type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

func (a *T850) BangBang() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}
func (r *R2D2) BangBang() error {
	return nil
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	t := T850{
		Name: "The Terminator",
	}
	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on !")
	}

	err = Boot(&t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on !")
	}
}

//Go语言没有提供类和类继承等面向对象功能，但结构体和方法集弥补了这部分不足，提供了一些面向对象元素。
//由此可以说Go在不使用类和继承的情况下提供了类似于面向对象编程的功能，虽然这一点还存在争议。
//在Go语言中，接口以声明的方式提供了多态，因为接口描述了必须提供的方法集以及这些方法的函数签名。
//接口实现可以包含其他方法（BangBang）
