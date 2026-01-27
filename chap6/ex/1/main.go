package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileSearch(dir string, filename string, wg *sync.WaitGroup, mx *sync.Mutex, output *[]string) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			mx.Lock()
			*output = append(*output, fpath)
			mx.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg, mx, output)
		}
	}
	wg.Done()
}

func main() {
	out := make([]string, 0)
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &wg, &mx, &out)
	wg.Wait()
	mx.Lock()
	fmt.Println(strings.Join(out, "\n"))
	mx.Unlock()
}
