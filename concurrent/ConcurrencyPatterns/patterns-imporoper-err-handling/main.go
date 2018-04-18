package main

import (
	"net/http"
	"fmt"
)

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response{
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil{
					fmt.Println(err)
					/*
					Here we see the goroutine doing its best to signal that there's an error.
					1. What else can it do?
					2. It can't pass it back!
					3. How many errors is too many?
					4. Does it continue making requests?
					 */
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...){
		fmt.Printf("Response: %v\n", response.Status)
	}
}
