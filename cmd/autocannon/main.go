package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func request(target string) {
	res, err := http.Get(target)
    if err != nil {
        panic(err)}

	fmt.Println("status", res.Status)

	markDone()
}

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
	res, err := http.Get("http://localhost:3000")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// respChan <- res

	doneChan <- res.Status
	markDone()
}

func main() {
	var start = time.Now()


	// num inflight
	// total rps

	var doneChan = make(chan string)
	for range 1900 {
		elapsed := time.Since(start).Seconds()
		go do(doneChan)

		// if len(doneChan) > 100 {
		// 	continue
		// }
		
		rps := float64(readDone())/elapsed
		fmt.Printf("rps: %f elapsed: %f", rps, elapsed)
		select {
		case <-doneChan:
			fmt.Printf(<-doneChan)
		default:
			continue
		}
		fmt.Println()
	}
}
