package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// Adi

func main() {
	var HEIGHT, WIDTH int
	fmt.Scanf("%d %d", &HEIGHT, &WIDTH)

	full_map := make([][]int, HEIGHT)

	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			fmt.Scanf("%d", &full_map[i][j])
		}
	}

	for i := 0; i < 7*WIDTH; i++ {
		if i == 0 || i == 7*WIDTH-1 {
			ap.PutRune(' ')
			if i == 7*WIDTH-1 {
				ap.PutRune('\n')
			}
		} else {
			ap.PutRune('_')
		}
	}

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if j >= WIDTH-1 {
				printCell(full_map[i][j], true)
			} else {
				printCell(full_map[i][j], false)
			}
		}
	}
}

func printFullRow(value rune, isLast bool) {
	if !isLast {
		for i := 0; i < 7; i++ {
			if i == 0 {
				ap.PutRune('|')
			} else {
				ap.PutRune(value)
			}
		}
	} else {
		for i := 0; i < 7; i++ {
			if i == 0 || i == 6 {
				ap.PutRune('|')
			} else {
				ap.PutRune(value)
			}
		}
	}
}

func printRuneAtCentre(value rune, isLast bool) {
	if !isLast {
		for i := 0; i < 7; i++ {
			if i == 0 {
				ap.PutRune('|')
			} else if i == 3 {
				ap.PutRune(value)
			} else {
				ap.PutRune(' ')
			}
		}
	} else {
		for i := 0; i < 7; i++ {
			if i == 0 || i == 6 {
				ap.PutRune('|')
			} else if i == 3 {
				ap.PutRune(value)
			} else {
				ap.PutRune(' ')
			}
		}
	}
}

// Azat
func printCell(value int, isLast bool) {
	if value == 0 {
		for i := 0; i < 3; i++ {

			printFullRow('X', isLast)

			ap.PutRune('\n')
		}
	} else if value == 1 {
		for i := 0; i < 3; i++ {
			if i == 2 {
				printFullRow('_', isLast)
			} else {
				printFullRow(' ', isLast)
			}
			ap.PutRune('\n')
		}
	} else if value == 2 {
		for i := 0; i < 3; i++ {

			if i == 1 {
				printRuneAtCentre('>', isLast)
			} else if i == 2 {
				printFullRow('_', isLast)
			} else {
				printFullRow(' ', isLast)
			}
			ap.PutRune('\n')
		}
	} else if value == 3 {
		for i := 0; i < 3; i++ {

			if i == 1 {
				printRuneAtCentre('@', isLast)
			} else if i == 2 {
				printFullRow('_', isLast)
			} else {
				printFullRow(' ', isLast)
			}
			ap.PutRune('\n')
		}
	}
}

// Additional tasks Yerkanat
