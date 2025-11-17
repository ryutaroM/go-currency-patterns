package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func countLetters(url string, frequency map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Server returning error status code" + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)

	text := strings.ToLower(string(body))
	regex := regexp.MustCompile(`\b[a-z]+\b`)
	words := regex.FindAllString(text, -1)
	for _, word := range words {
		frequency[word] += 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int, 0)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency) //happens DATA RACE because multiple goroutines access frequency map concurrently
	}
	time.Sleep(10 * time.Second)
	for c, count := range frequency {
		fmt.Println(c, ":->", count)
	}
}
