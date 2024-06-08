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

	horiz_coord := make([][7]rune, WIDTH)
	horiz_coord_count(WIDTH, &horiz_coord)

	full_map := make([][]int, HEIGHT)

	// Read input grid values
	var cell rune
	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			fmt.Scanf("%c", &cell)
			full_map[i][j] = int(cell - rune('0'))
		}
		fmt.Scanf("%c", &cell)
	}

	fmt.Scanf("%c%c%c", &wall, &player, &award)

	// Call printMap function to print the map
	printMap(HEIGHT, WIDTH, full_map, &horiz_coord)
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
func printCell(value, y, x int) {
	if value == 0 {
		colorize(value)
		ap.PutRune(wall)
		colorize(4)
	} else if value == 1 {
		if y == 0 {
			colorize(value)
			ap.PutRune('_')
			colorize(4)
		} else {
			colorize(value)
			ap.PutRune(' ')
			colorize(4)
		}
	} else if value == 2 {
		if y == 2 && x == 4 {
			colorize(value)
			ap.PutRune(player)
			colorize(4)
		} else if y == 0 {
			colorize(value)
			ap.PutRune('_')
			colorize(4)
		} else {
			colorize(value)
			ap.PutRune(' ')
			colorize(4)
		}
	} else if value == 3 {
		if y == 2 && x == 4 {
			colorize(value)
			ap.PutRune(award)
			colorize(4)
		} else if y == 0 {
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

func printMap(height, width int, value [][]int, horiz_coord *[][7]rune) {
	// Loop to iterate over every row.
	// Multiplied by 3 because every row has 3 rows inside.
	// And additional 1 to print first line
	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 {
			ap.PutRune(' ')
			ap.PutRune(rune(((h + 1) / 3) + '0'))
			ap.PutRune(' ')
		} else {
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune(' ')
		}

		// Loop to iterate over every column
		// Multiplied by 8 because every column has 7 runes inside and borders
		if h == 0 {
			for w := 0; w < (8*width)+1; w++ {
				if w%8 != 0 {
					ap.PutRune((*horiz_coord)[w/8][w%8-1])
				} else {
					ap.PutRune(' ')
				}
			}
			ap.PutRune('\n')
		}

		for w := 0; w < (8*width)+1; w++ {
			if w%8 == 0 && h != 0 {
				ap.PutRune('|')
			} else if h == 0 {
				if w == 0 {
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune(' ')
				}
				if w != 0 && w != 8*width {
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
}

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
