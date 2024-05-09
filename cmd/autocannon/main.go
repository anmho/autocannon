package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client http.Client



var doneMu = sync.RWMutex{}
var done = 0

func markDone() {
	doneMu.Lock()
	defer doneMu.Unlock()
	done++
}

func readDone() int {
	doneMu.RLock()
	defer doneMu.RUnlock()
	return done
}


func do(doneChan chan string) {
	
	res, err := client.Get("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	// respChan <- res

	doneChan <- res.Status
	
	markDone()
}

func main() {
	var start = time.Now()

	client = http.Client{
		Transport: &http.Transport{MaxConnsPerHost: 100},
	}


	total := 100_000
	var doneChan = make(chan string)
	for range total {
		go do(doneChan)

		
		select {
		case <-doneChan:
			fmt.Println(<-doneChan)
		default:
			continue
		}		
	}
	elapsed := time.Since(start).Seconds()
	rps := float64(total) / elapsed
	fmt.Printf("rps: %f elapsed: %f", rps, elapsed)

}
