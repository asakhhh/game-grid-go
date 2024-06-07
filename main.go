package main

import (
	"fmt"

	"github.com/alem-platform/ap"
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
	printCell(WIDTH)
}

func printFullRow(value rune) {
	for i := 0; i < 7; i++ {
		if i == 0 {
			ap.PutRune('|')
		} else {
			ap.PutRune(value)
		}
	}
}

func printRuneAtCentre(value rune) {
	for i := 0; i < 7; i++ {
		if i == 0 {
			ap.PutRune('|')
		} else if i == 3 {
			ap.PutRune(value)
		} else {
			ap.PutRune(' ')
		}
	}
}

// Azat
func printCell(value int) {
	ap.PutRune(' ')
	// for i := 0; i < 7; i++ {
	// 	ap.PutRune('_')
	// }
	ap.PutRune('\n')

	if value == 0 {
		for i := 0; i < 3; i++ {

			printFullRow('X')

			ap.PutRune('\n')
		}
	} else if value == 1 {
		for i := 0; i < 3; i++ {

			if i == 2 {
				printFullRow('_')
			} else {
				printFullRow(' ')
			}
			ap.PutRune('\n')
		}
	} else if value == 2 {
		for i := 0; i < 3; i++ {

			if i == 1 {
				printRuneAtCentre('>')
			} else if i == 2 {
				printFullRow('_')
			} else {
				printFullRow(' ')
			}
			ap.PutRune('\n')
		}
	} else if value == 3 {
		for i := 0; i < 3; i++ {

			if i == 1 {
				printRuneAtCentre('@')
			} else if i == 2 {
				printFullRow('_')
			} else {
				printFullRow(' ')
			}
			ap.PutRune('\n')
		}
	}
}

// Additional tasks Yerkanat
