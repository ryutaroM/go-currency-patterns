package main

import (
	"fmt"
	"time"
)

const (
	passwordGuess = "go far"
	alpahabet     = "abcdefghijklmnopqrstuvwxyz "
)

func toBase27(n int) string {
	result := ""
	for n > 0 {
		result = string(alpahabet[n%27]) + result
		n /= 27
	}
	return result
}

func guessPassword(from int, upto int, stop chan int, result chan string) {
	for guessN := from; guessN < upto; guessN++ {
		select {
		case <-stop:
			fmt.Printf("Stopped at %d [%d, %d]\n", guessN, from, upto)
			return
		default:
			if toBase27(guessN) == passwordGuess {
				result <- toBase27(guessN)
				close(stop)
				return
			}
		}
	}
}

func main() {
	finished := make(chan int)
	passwordFound := make(chan string)

	for i := 1; i <= 387_420_488; i += 10_000_000 {
		go guessPassword(i, i+10_000_000, finished, passwordFound)
	}

	fmt.Println("password found:", <-passwordFound)
	close(passwordFound)
	time.Sleep(5 * time.Second)
}
