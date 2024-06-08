package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// colors of cells
var (
	RESET      = []rune{'\033', '[', '0', 'm'}
	RED        = []rune{'\033', '[', '3', '1', 'm'}
	WHITE      = []rune{'\033', '[', '9', '7', 'm'}
	BLUE       = []rune{'\033', '[', '3', '4', 'm'}
	YELLOW     = []rune{'\033', '[', '3', '3', 'm'}
	colorCodes = [][]rune{RED, WHITE, BLUE, YELLOW, RESET}
)

// Change symbols of bonus task
var (
	wall   = '0'
	player = '2'
	award  = '3'
)

// function that colorizes cells on the given input value
func colorize(val int) {
	if val >= 0 && val < len(colorCodes) {
		for _, r := range colorCodes[val] {
			ap.PutRune(r)
		}
	}
}

// printCell function prints a single cell based on its value
func printCell(value, y, x int) {
	colorize(value)
	if value == 0 {
		ap.PutRune(wall)
	} else if value == 1 {
		if y == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(' ')
		}
	} else if value == 2 {
		if y == 2 && x == 4 {
			ap.PutRune(player)
		} else if y == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(' ')
		}
	} else if value == 3 {
		if y == 2 && x == 4 {
			ap.PutRune(award)
		} else if y == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(' ')
		}
	}
	colorize(4)
}

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

// function that returns result of logroithm on base of N
func logofN(height int, n int) int {
	length := 1

	for ; height > 0; height /= n {
		length++
	}

	return length
}

// Loop to iterate over every column
// Multiplied by 8 because every column has 7 runes inside and borders
func printColumn(width int, horiz_coord *[][7]rune) {
	for w := 0; w < (8*width)+1; w++ {
		if w%8 != 0 {
			ap.PutRune((*horiz_coord)[w/8][w%8-1])
		} else {
			ap.PutRune(' ')
		}
	}
	ap.PutRune('\n')
}

func printRow(width int, length int, value [][]int, h int) {
	for w := 0; w < (8*width)+1; w++ {
		if w%8 == 0 && h > 0 {
			ap.PutRune('|')
		} else if h == 0 {
			if w == 0 {
				for i := 0; i < length+2; i++ {
					ap.PutRune(' ')
				}
			}
			if w > 0 && w != 8*width {
				ap.PutRune('_')
			} else {
				ap.PutRune(' ')
			}
		} else {
			printCell(value[(h-1)/3][w/8], h%3, w%8) // Print cell content
		}
	}
	ap.PutRune('\n')
}

// themost important function that uses all other functions to print map
func printMap(height, width int, value [][]int, horiz_coord *[][7]rune) {
	// Loop to iterate over every row.
	// Multiplied by 3 because every row has 3 rows inside.
	// And additional 1 to print first line

	length := logofN(height, 10)

	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 {
			printNumber((h+1)/3, length)
		} else {
			for i := 0; i < length+2; i++ {
				ap.PutRune(' ')
			}
		}
		// print horizontal coordinates
		if h == 0 {
			printColumn(width, horiz_coord)
		}
		// print
		printRow(width, length, value, h)
	}
}

// helper function to add character when "Z" on horizontal axis is reached. It starts from "AA" 'AB' ...
func add_one(arr *[][7]rune, i, j int) {
	if (*arr)[i][j] == 'Z' {
		(*arr)[i][j] = 'A'
		if j > 0 {
			add_one(arr, i, j-1)
		}
	} else if (*arr)[i][j] == ' ' {
		(*arr)[i][j] = 'A'
	} else {
		(*arr)[i][j] = rune((*arr)[i][j] + 1)
	}
}

// function prints horizontal coordinates of map
func horiz_coord_count(width int, arr *[][7]rune) {
	for i := 0; i < 6; i++ {
		(*arr)[0][i] = ' '
	}
	(*arr)[0][6] = 'A'

	for i := 1; i < width; i++ {
		for j := 0; j < 7; j++ {
			(*arr)[i][j] = (*arr)[i-1][j]
		}
		add_one(arr, i, 6)
	}

	for i := 0; i < width; i++ {
		if (*arr)[i][5] == ' ' {
			(*arr)[i][3] = (*arr)[i][6]
			(*arr)[i][6] = ' '
		} else if (*arr)[i][4] == ' ' {
			(*arr)[i][3] = (*arr)[i][5]
			(*arr)[i][4] = (*arr)[i][6]
			(*arr)[i][5] = ' '
			(*arr)[i][6] = ' '
		} else if (*arr)[i][3] == ' ' {
			(*arr)[i][2] = (*arr)[i][4]
			(*arr)[i][3] = (*arr)[i][5]
			(*arr)[i][4] = (*arr)[i][6]
			(*arr)[i][5] = ' '
			(*arr)[i][6] = ' '
		} else if (*arr)[i][2] == ' ' {
			(*arr)[i][2] = (*arr)[i][3]
			(*arr)[i][3] = (*arr)[i][4]
			(*arr)[i][4] = (*arr)[i][5]
			(*arr)[i][5] = (*arr)[i][6]
			(*arr)[i][6] = ' '
		} else if (*arr)[i][1] == ' ' {
			(*arr)[i][1] = (*arr)[i][2]
			(*arr)[i][2] = (*arr)[i][3]
			(*arr)[i][3] = (*arr)[i][4]
			(*arr)[i][4] = (*arr)[i][5]
			(*arr)[i][5] = (*arr)[i][6]
			(*arr)[i][6] = ' '
		}
	}
}

// function scans widht and height of map. Retuns scanned values
func readWidthAndHeight() (int, int, [][7]rune) {
	var width, height int

	if _, err := fmt.Scanf("%d %d", &height, &width); err != nil {
		return -1, -1, nil
	}

	// make needed to implement Map numbers on x and y-axises.
	return height, width, make([][7]rune, width)
}

// function scans h lines each with w number of characters
// also it scans for Change symbols
func readGridValues(WIDTH int, HEIGHT int) [][]int {
	var cell rune
	full_map := make([][]int, HEIGHT)

	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			if _, err := fmt.Scanf("%c", &cell); err != nil {
				return nil
			}
			if full_map[i][j] = int(cell - rune('0')); full_map[i][j] < 0 || full_map[i][j] > 3 {
				return nil
			}

		}
		if _, err := fmt.Scanf("%c", &cell); err != nil {
			return nil
		}
	}
	if _, err := fmt.Scanf("%c%c%c", &wall, &player, &award); err != nil {
		return nil
	}

	return full_map
}

func printString(str string) {
	for i := 0; i < len(str); i++ {
		ap.PutRune(rune(str[i]))
	}
}
