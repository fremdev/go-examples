package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://bing.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Site is not available:", link)
		c <- "Might be down"
		return
	}

	fmt.Println("Site is OK:", link)
	c <- "Site is OK"
}
