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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel to communicate between goroutines
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// this will block until a value is sent to the channel
	// fmt.Println(<-c)

	// for i:= 0; i<len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// this will block until all links are checked
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // send the link back to the channel for further checking
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}