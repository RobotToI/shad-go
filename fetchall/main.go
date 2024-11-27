//go:build !solution

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: urlfetch <link1> <link2> ...")
	}

	start := time.Now()

	var urls []string = os.Args[1:]
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading %s: %v\n", url, err)
		return
	}
	ch <- fmt.Sprintf("%.3fs %s", time.Since(start).Seconds(), url)
}
