package main

import (
	"net/http"
	"log"
	"context"
	"time"
	"net"
	"io"
	"os"
)

func main() {
	req, err := http.NewRequest("GET","https://www.ardanlabs.com/blog/post/index.xml", nil)
	if err != nil{
		log.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()
	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns: 100,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := http.Client{
		Transport: &tr,
	}
	// Make the web call in a separate goroutine so it can be cancelled.
	ch := make(chan error, 1)
	go func() {
		log.Println("Starting Request")
		// Make the web call and return any error.
		resp, err := client.Do(req)
		if err != nil{
			ch <- err
			return
		}
		// Close the response body on the return.
		defer resp.Body.Close()
		// Write the response to stdout
		io.Copy(os.Stdout, resp.Body)
		ch <- nil
	}()

	// Wait the request or timeout.
	select {
	case <- ctx.Done():tr.CancelRequest(req)
		log.Println("timeout, cancel work ...")
		tr.CancelRequest(req)
		log.Println(<-ch)
	case err := <-ch:
		if err != nil{
			log.Println(err)
		}
	}
}
