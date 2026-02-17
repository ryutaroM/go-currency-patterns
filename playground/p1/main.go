package main

import (
	"fmt"
)

func f() int {
	return -1
}

func main() {
	ch := make(chan int, 1)
	go func() {
		fmt.Println("go func called")
	}()
	ch <- f()
	fmt.Println(<-ch)
}
