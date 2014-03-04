package main

import (
	"fmt"
)

func main() {
	fmt.Print("aa")

	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:2]
	fmt.Print(arr)
	slice = append(slice, 6, 7)
	fmt.Print(arr)
	fmt.Print(slice)
}
