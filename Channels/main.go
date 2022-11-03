package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.co.in",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "migth not up")
		c <- "migth not up"
		return
	}
	c <- link + " Working!"
}
