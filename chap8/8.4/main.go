package main

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
