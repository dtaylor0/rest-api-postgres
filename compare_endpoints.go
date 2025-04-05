package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	REQ_COUNT = 10000
)

type Stats struct {
	minRTT     int64
	maxRTT     int64
	aggRTT     int64
	okCount    int
	errorCount int
}

type TestConfig struct {
	name     string
	url      string
	requests int
}

func main() {
	runPerfTests([]TestConfig{
		{name: "typescript", url: "http://localhost:3000/api/msft", requests: REQ_COUNT},
		{name: "golang", url: "http://localhost:8080/api/msft", requests: REQ_COUNT},
	})
}

func runPerfTests(tests []TestConfig) {
	var stats string
	for _, test := range tests {
		fmt.Printf("Name: [%s] Requests: [%d] URL: [%s]\n", test.name, test.requests, test.url)
		fmt.Println("Running performance test...")
		stats = runPerfTest(test.url, test.requests)
		fmt.Println(stats)
		fmt.Println()
	}
}

func runPerfTest(endpoint string, requests int) string {
	wg := sync.WaitGroup{}
	mutex := &sync.Mutex{}
	wg.Add(requests)
	stats := Stats{}
	for range requests {
		go func() {
			defer wg.Done()

			// Create http client
			client := http.Client{Timeout: time.Second * 20}

			// Time api response
			start := time.Now()
			res, err := client.Get(endpoint)
			end := time.Now()

			mutex.Lock()
			defer mutex.Unlock()
			// track error count separately from ok
			if err != nil || res.StatusCode >= 400 {
				stats.errorCount++
				return
			}

			// Update stats for the api requests
			stats.okCount++
			currRTT := end.Sub(start).Milliseconds()

			if stats.maxRTT < currRTT {
				stats.maxRTT = currRTT
			} else if stats.minRTT > currRTT || stats.minRTT == 0 {
				stats.minRTT = currRTT
			}

			stats.aggRTT += currRTT
		}()
	}
	wg.Wait()
	meanRTT := 0
	if stats.okCount > 0 {
		meanRTT = int(stats.aggRTT) / stats.okCount
	}
	return fmt.Sprintf(
		"Results\n=======\nOK: %d\nError: %d\nMin: %dms\nMax: %dms\nMean: %dms\n",
		stats.okCount,
		stats.errorCount,
		stats.minRTT,
		stats.maxRTT,
		meanRTT,
	)
}
