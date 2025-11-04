package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Less than 1 file.")
		return
	}

	ch := make(chan struct{}, 3)

	for _, f := range os.Args[1:] {
		go do(f, ch)

	}

	for i := 0; i < len(os.Args[1:]); i++ {
		<-ch
	}
}

func do(f string, ch chan struct{}) {
	defer func() {
		ch <- struct{}{}
	}()
	file, err := os.OpenFile(f, os.O_RDONLY, 0)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		os.Exit(2)
	}

}
