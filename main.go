package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	timeElapsed := time.Now()
	url := flag.String("url", "", "URL to call")
	requests := flag.Int("requests", 1, "How many requests to call. Default: 1")
	workers := flag.Int("concurrency", 1, "How many requests to call at same time. Default: 1")
	flag.Parse()

	validateArguments(url, requests, workers)
	times := calculateReqPerWorkers(requests, workers)

	report := make(map[int]int)
	doRequests(times, workers, url, report)

	printReport(report, timeElapsed)
}

func validateArguments(url *string, requests *int, workers *int) {
	if *url == "" {
		panic("Url argument is mandatory.")
	}
	if *requests < 1 {
		panic("Number of requests must be higher than 1")
	}
	if *workers < 1 {
		panic("Number of concurrency must be higher than 1")
	}
}

func calculateReqPerWorkers(requests *int, workers *int) int {
	times := *requests / *workers
	if times < 0 {
		times = 1
	}
	return times
}

func doRequests(times int, workers *int, url *string, report map[int]int) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		for worker := 0; worker < *workers; worker++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := http.Get(*url)
				if err != nil || resp == nil {
					fmt.Println(http.StatusInternalServerError)
					mutex.Lock()
					report[http.StatusInternalServerError]++
					mutex.Unlock()
				} else {
					fmt.Println(resp.StatusCode)
					mutex.Lock()
					report[resp.StatusCode]++
					mutex.Unlock()
				}
			}()
		}
	}
	wg.Wait()
}

func printReport(report map[int]int, timeElapsed time.Time) {
	fmt.Println("------------------------------")
	fmt.Println("--------Report Summary--------")
	for httpStatus, count := range report {
		fmt.Printf("Http code: %v - Count: %v\n", httpStatus, count)
	}
	fmt.Printf("Time elapsed: %v\n", time.Since(timeElapsed))
	fmt.Println("------------------------------")
}
