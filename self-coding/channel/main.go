package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(c, link)
	}

	for l := range c {
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(c, li)
		}(l)
	}
}

func checkLink(c chan string, link string) {
	time.Sleep(2 * time.Second)

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, " is ok!")
	c <- link
}
