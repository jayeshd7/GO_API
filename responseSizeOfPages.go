package main

import "fmt"

type Page struct {
	Url  string
	Size int
}

func responseSize(url string, channel chan Page) {
	channel <- Page{url, len(url)}
}

func mainResp() {
	pages := make(chan Page)
	urls := []string{"https://www.google.com/", "https://www.yahoo.com/", "https://www.bing.com/"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println("url:", page.Url, "size:", page.Size)
	}
}
