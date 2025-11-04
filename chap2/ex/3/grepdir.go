package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Less than 1 file.")
		return
	}

	dir := os.Args[2]

	entries, err := os.ReadDir(dir)
	if err != nil {
		os.Exit(2)
	}

	fileCount := 0

	ch := make(chan struct{}, 3)

	for _, f := range entries {
		fileCount++
		go do(os.Args[1], filepath.Join(dir, f.Name()), ch)
	}

	for i := 0; i < fileCount; i++ {
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
