package main

import (
	"fmt"
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() *http.Response {
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != 200 {
		log.Fatalln("FAILED : %d", res.StatusCode)
	}
	return res
}
