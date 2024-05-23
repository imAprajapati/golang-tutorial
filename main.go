package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // Pointer

func main() {
	websitelist := []string{
		"https://www.google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
		"https://youtube.com",
	}
	for _, endpoint := range websitelist {
		go getStatusCode(endpoint)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
