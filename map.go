package main

import (
	"github.com/alem-platform/ap"
)

// Default symbols that can be changed by user
var (
	wall    = 'X'
	empty   = ' '
	player  = '>'
	award   = '@'
	floor   = '_'
	symbols = []rune{wall, empty, player, award, floor}
)

// function printing map of w*h
func printMap(height, width int, value [][]int) {
	horiz_coord := make([][7]rune, width)
	horiz_coord_count(width, &horiz_coord)
	length := logofN(height, 10)

	// Loop to iterate over every row.
	// Multiplied by 3 as every row has 3 rows inside.
	// And additional 1 to print first line

	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 {
			printVerticalCoord((h+1)/3, length)
		} else {
			for i := 0; i < length+2; i++ {
				ap.PutRune(' ')
			}
		}
		if h == 0 {
			printHorizntalCoord(width, &horiz_coord)
		}
		// prints row of the map
		printRow(width, length, value, h)
	}
}

// helper function that prints one row of a map
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

// printCell function prints a single cell of a row based on its value
func printCell(value, y, x int) {
	colorizeBG(value)
	colorizeFG(value)

	if value == 0 { // print value of change symbol for wall
		ap.PutRune(symbols[value])
	} else if value == 1 { // print empty cell
		if y == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(symbols[value])
		}
	} else { //  print value of change symbol
		if y == 2 && x == 4 {
			ap.PutRune(symbols[value])
		} else if y == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(symbols[1])
		}
	}
	colorizeBG(1)
	colorizeFG(1)
}

// Loop to iterate over every column
// Multiplied by 8 because every column has 7 runes inside and borders
func printHorizntalCoord(width int, horiz_coord *[][7]rune) {
	for w := 0; w < (8*width)+1; w++ {
		if w%8 != 0 {
			ap.PutRune((*horiz_coord)[w/8][w%8-1])
		} else {
			ap.PutRune(' ')
		}
	}
	ap.PutRune('\n')
}

// helper function to add character when "Z" on horizontal axis is reached. It starts from "AA" 'AB' ...
func nextLetterAfterZ(arr *[][7]rune, i, j int) {
	if (*arr)[i][j] == 'Z' {
		(*arr)[i][j] = 'A'
		if j > 0 {
			nextLetterAfterZ(arr, i, j-1)
		}
	} else if (*arr)[i][j] == ' ' {
		(*arr)[i][j] = 'A'
	} else {
		(*arr)[i][j] = rune((*arr)[i][j] + 1)
	}
}

// function calculates horizontal coordinates of map
func horiz_coord_count(width int, arr *[][7]rune) {
	for i := 0; i < 6; i++ {
		(*arr)[0][i] = ' '
	}
	(*arr)[0][6] = 'A'

	for i := 1; i < width; i++ {
		for j := 0; j < 7; j++ {
			(*arr)[i][j] = (*arr)[i-1][j]
		}
		nextLetterAfterZ(arr, i, 6)
	}
	alignByCentre(width, arr)
}

// function that aligns by centre horizontal notations of the map
func alignByCentre(width int, arr *[][7]rune) {
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
