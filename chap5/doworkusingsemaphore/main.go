package main

import (
	"fmt"
	semaphore "go-concurrency-patterns/chap5/semaphore"
)

func main() {
	semaphore := semaphore.NewSemaphore(0)
	for i := 0; i < 50000; i++ {
		go dowork(semaphore)
		fmt.Println("Waiting for child goroutine")
		semaphore.Acquire()
		fmt.Println("Child goroutine finished")
	}
}

func dowork(semaphore *semaphore.Semaphore) {
	fmt.Println("Work Started")
	fmt.Println("Work Finished")
	semaphore.Release()
}
