package main

import (
	"github.com/alem-platform/ap"
	"fmt"
)

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	WHITE  = "\033[97m"
	BLUE   = "\033[34m"
	YELLOW = "\033[33m"
)
// Adi
func main() {
	var HEIGHT int
	var WIDTH int

	fmt.Scanf("%d", &HEIGHT)
	fmt.Scanf("%d", &WIDTH)
}

// Azat 
func printCell(value rune) {
	for i := 0; i < 7; i++ {
		ap.PutRune(value)
	}
}

// Additional tasks Yerkanat