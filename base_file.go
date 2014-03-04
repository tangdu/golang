// base_file
package main

import (
	"fmt"
	"log"
	"os"
)

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		fmt.Print(buf[:n])
	}
}

func main() {
	fmt.Println("Hello World!")
	readFile("e:/1.html")
}
