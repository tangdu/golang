package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Host)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["name33"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintln(w, "你好,h")
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server error", err)
	}
}
