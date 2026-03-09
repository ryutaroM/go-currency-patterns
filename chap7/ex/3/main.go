package main

import (
	"fmt"
	"math/rand"
)

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	resultCh := make(chan []int)
	for i := 0; i < 10; i++ {
		go func() {
			resultCh <- findFactors(rand.Intn(1000000))
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-resultCh)
	}
}
