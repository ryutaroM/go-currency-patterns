package main

import (
	"fmt"
	origin "go-concurrency-patterns/chap4/originalreaderwriter"
	"time"
)

func main() {
	rwMutex := origin.ReaderWriterMutex{}
	for i := 0; i < 2; i++ {
		go func() {
			for {
				rwMutex.ReadLock()
				time.Sleep(1 * time.Second)
				fmt.Println("Read done")
				rwMutex.ReadUnlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	rwMutex.WriteLock()
	fmt.Println("Write finished")
}
