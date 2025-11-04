package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Less than 1 file.")
		return
	}

	ch := make(chan struct{}, 3)

	for _, f := range os.Args[2:] {
		go do(os.Args[1], f, ch)

	}

	for i := 0; i < len(os.Args[2:]); i++ {
		<-ch
	}
}

func do(haystack, f string, ch chan struct{}) {
	defer func() {
		ch <- struct{}{}
	}()

	file, err := os.OpenFile(f, os.O_RDONLY, 0)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), haystack) {
			fmt.Printf("filename:%s has %q\n", f, haystack)
		}
	}
}
