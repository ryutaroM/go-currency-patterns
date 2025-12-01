package main

import (
	"fmt"
	"sync"
	"time"
)

func countDown(second *int, mux *sync.Mutex) {
	mux.Lock()
	cp := *second
	mux.Unlock()
	for cp > 0 {
		time.Sleep(1 * time.Second)
		mux.Lock()
		*second--
		cp = *second
		mux.Unlock()
	}
}

func main() {
	count := 5
	mux := sync.Mutex{}
	cp := count
	go countDown(&count, &mux)
	for cp > 0 {
		time.Sleep(500 * time.Millisecond)
		mux.Lock()
		fmt.Println(count)
		cp = count
		mux.Unlock()
	}
}
