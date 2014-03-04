package main

import (
	"fmt"
	"os"
)

func add(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

func mod(a int, b int) (c int, d int) {
	return a, b
}

func testDefer() {
	file, err := os.Open("hello.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

}

/*
Ｃ语言中提供了地址运算符&来表示变量的地址。其一般形式为： & 变量名； 如&a变示变量a的地址，&b表示变量b的地址。
变量本身必须预先说明。设有指向整型变量的指针变量p，如要把整型变量a 的地址赋予p可以有以下两种方式：
(1)指针变量初始化的方法 int a;
int *p=&a;
(2)赋值语句的方法 int a;
int *p;
p=&a;
不允许把一个数赋予指针变量，故下面的赋值是错误的： int *p;p=1000; 被赋值的指针变量前不能再加“*”说明符，如写为*p=&a 也是错误的
-------------------------
• Go语言中string，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
*/
func main() {
	a := 1
	b := add(a)

	fmt.Println(b, a)

	c := 1
	d := add2(&c)
	fmt.Println(d, c)

	fmt.Println(mod(1, 3))
	//testDefer()
}
