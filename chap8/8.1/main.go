package main

import (
	"fmt"
	"time"
)

func writeEvery(msg string, seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		for {
			time.Sleep(seconds)
			messages <- msg
		}
	}()
	return messages
}

func main() {
	messagesFromA := writeEvery("Tick", 1*time.Second)
	messagesFromB := writeEvery("Tock", 3*time.Second)
	for {
		select {
		case msg := <-messagesFromA:
			fmt.Println(msg)
		case msg := <-messagesFromB:
			fmt.Println(msg)
		}
	}
}
