package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

var (
	RESET  = []rune{'\033', '[', '0', 'm'}
	RED    = []rune{'\033', '[', '3', '1', 'm'}
	WHITE  = []rune{'\033', '[', '9', '7', 'm'}
	BLUE   = []rune{'\033', '[', '3', '4', 'm'}
	YELLOW = []rune{'\033', '[', '3', '3', 'm'}
)

var (
	wall   = '0'
	player = '2'
	award  = '3'
)

func main() {
	var HEIGHT, WIDTH int
	fmt.Scanf("%d %d", &HEIGHT, &WIDTH)

	full_map := make([][]int, HEIGHT)

	// Read input grid values
	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			fmt.Scanf("%d", &full_map[i][j])
		}
	}

	fmt.Scanf("%c%c%c", &wall, &player, &award)

	// Call printMap function to print the map
	printMap(HEIGHT, WIDTH, full_map)
}

func colorize(val int) {
	if val == 0 {
		for _, r := range RED {
			ap.PutRune(r)
		}
	} else if val == 2 {
		for _, r := range BLUE {
			ap.PutRune(r)
		}
	} else if val == 3 {
		for _, r := range YELLOW {
			ap.PutRune(r)
		}
	} else if val == 4 {
		for _, r := range RESET {
			ap.PutRune(r)
		}
	}
}

// printCell function prints a single cell based on its value
func printCell(value int, x int, y int) {
	if value == 0 {
		colorize(value)
		ap.PutRune(wall)
		colorize(4)
	} else if value == 1 {
		if x%3 == 0 {
			colorize(value)
			ap.PutRune('_')
			colorize(4)
		} else {
			colorize(value)
			ap.PutRune(' ')
			colorize(4)
		}
	} else if value == 2 {
		if x == 2 && y == 4 {
			colorize(value)
			ap.PutRune(player)
			colorize(4)
		} else if x == 0 {
			colorize(value)
			ap.PutRune(' ')
			colorize(4)
		} else {
			colorize(value)
			ap.PutRune(' ')
			colorize(4)
		}
	} else if value == 3 {
		if x == 2 && y == 4 {
			colorize(value)
			ap.PutRune(award)
			colorize(4)
		} else if x == 0 {
			colorize(value)
			ap.PutRune('_')
			colorize(4)
		} else {
			colorize(value)
			ap.PutRune(' ')
			colorize(4)
		}
	}
}

func printMap(height int, width int, value [][]int) {
	// Loop to iterate over every row.
	// Multiplied by 3 because every row has 3 rows inside.
	// And additional 1 to print first line
	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 && h != 0 {
			ap.PutRune(' ')
			ap.PutRune(rune(((h + 1) / 3) + '0'))
		} else {
			ap.PutRune(' ')
			ap.PutRune(' ')
		}

		// Loop to iterate over every column
		// Multiplied by 8 because every column has 7 runes inside and borders
		if h == 0 {
			for w := 0; w < (8*width)+1; w++ {
				if (w+3)%7 == 0 && w != 0 {
					ap.PutRune(rune(64 + (w+3)/7)) // Print horizontal coordinates
				} else {
					ap.PutRune(' ')
				}
			}
			ap.PutRune('\n')
		}

		for w := 0; w < (8*width)+1; w++ {
			if w == 0 && h != 0 || w%8 == 0 && h != 0 {
				ap.PutRune('|')
			} else if h == 0 {
				if w == 0 {
					ap.PutRune(' ')
					ap.PutRune(' ')
				}
				ap.PutRune('_')
			} else if h%3 != 0 {
				printCell(value[h/3][w/8], h%3, w%8) // Print cell content
			} else {
				printCell(value[h/3-1][w/8], h%3, w%8) // Print cell content
			}
		}
		ap.PutRune('\n')
	}
}
