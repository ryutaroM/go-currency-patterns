package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPU cores:", runtime.NumCPU())

	fmt.Println("GOMAXPROCS value:", runtime.GOMAXPROCS(0))
}
