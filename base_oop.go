package main

import (
	"fmt"
	"log"
	"os"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

func (s *Human) SayHi() { //继承
	fmt.Println("...human")
}

func (s *Student) SayHi() { //重写
	fmt.Println("...student")
}

func OpenFile() {
	file, err := os.Open("d.xtx")
	defer file.Close()
	if err != nil {
		log.Fatal("error is ", err)
	}

}

/*
*前面一章我们学习了字段的继承，那么你也会发现Go的一个神奇之处，
*method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。让我们来看下面这个例子
 */
func main() {
	mark := Student{Human{name: "相关", age: 23}, "一中"}
	mark.SayHi()

	OpenFile()
}
