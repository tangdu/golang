package main

import (
	"fmt"
)

//变量类型可以浑身是省写
const (
	a string = "唐杜"
	b int    = 2
)

type Student struct {
}

type GoFace interface{}

var c string = "dd" //全局变量,不使用不会报错

func main() {
	//c := "d" 区别? 1.此种方式不用会报错：c declared and not used 2.仅在函数内使用
	c = "d"
	//var d = "dd" //不使用会报错
	fmt.Print("..." + a + "-" + c)
	fmt.Print("ddd")
}
