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

func main() {
	var HEIGHT int
	var WIDTH int

	fmt.Scanf("%d", &HEIGHT)
	fmt.Scanf("%d", &WIDTH)
	printMap(HEIGHT, WIDTH)
}

func printCell(value rune) {
	for i := 0; i < 7; i++ {
		ap.PutRune(value)
	}
}

func printMap(height int, width int) {
	// Loop to iterate over every row.
	// Multiplied by 3 because every row has 3 rows inside.
	// And additional 1 to print first line
	for h := 0; h < (3*height)+1; h++ {
		// Loop to iterate over every column

		for w := 0; w < (7*width)+width+1; w++ {
			if w == 0 && h != 0 || w%8 == 0 && h != 0 {
				fmt.Print("|")
			} else if h == 0 || h%3 == 0 {
				fmt.Print("_")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
