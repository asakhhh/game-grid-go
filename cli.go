package main

import "github.com/alem-platform/ap"

// function that prints number using ap.PutRune
func printNumber(num, length int) {
	p, curlength := 1, 1

	for ; p*10 <= num; p *= 10 {
		curlength++
	}

	for i := 0; i <= length-curlength; i++ {
		ap.PutRune(' ')
	}

	for ; p > 0; p /= 10 {
		ap.PutRune(rune('0' + (num/p)%10))
	}
	ap.PutRune(' ')
}

// function that prints strings using ap.PutRune
func printString(str string) {
	for i := 0; i < len(str); i++ {
		ap.PutRune(rune(str[i]))
	}
}
