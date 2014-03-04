package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	url := "http://www.iteye.com/"
	t := time.Now()
	for i := 0; i < 100; i++ {
		craw(url)
	}

	fmt.Println()
	fmt.Print(time.Now().Sub(t))
}

func craw(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1573.2 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	//fmt.Print("in.........")
	if resp.StatusCode == 200 {
		//fmt.Print("success")

		respstrem, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
		html := string(respstrem)

		file, err := os.Create("f:/tmp/" + strconv.FormatInt(time.Now().UnixNano(), 3) + ".html")

		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(html)
		defer file.Close()
	}
}
