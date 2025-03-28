package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// colors for characters
var (
	RESET  = []rune{'\033', '[', '0', 'm'}
	RED    = []rune{'\033', '[', '3', '1', 'm'}
	WHITE  = []rune{'\033', '[', '9', '7', 'm'}
	BLUE   = []rune{'\033', '[', '3', '4', 'm'}
	YELLOW = []rune{'\033', '[', '3', '3', 'm'}

	colorCodes = [][]rune{RED, RESET, BLUE, YELLOW}
)

// colors for background
var (
	BgRed    = []rune("\033[41m")
	BgGreen  = []rune("\033[42m")
	BgYellow = []rune("\033[1;33;43m")
	BgBlue   = []rune("\033[44m")
	BgPurple = []rune("\033[45m")
	BgCyan   = []rune("\033[46m")
	BgWhite  = []rune("\033[47m")

	colorCodesBG = [][]rune{BgRed, RESET, BgBlue, BgYellow, BgPurple, BgCyan}
)

// function that colorizes Foreground
func colorizeFG(val int) {
	if val >= 0 && val < len(colorCodes) {
		for _, r := range colorCodes[val] {
			ap.PutRune(r)
		}
	}
}

// function that colorizes background
func colorizeBG(val int) {
	if val >= 0 && val < len(colorCodesBG) {
		for _, r := range colorCodesBG[val] {
			ap.PutRune(r)
		}
	}
}

// function that vaertical axis of map
func printVerticalCoord(num, length int) {
	p, curlength := 1, 1

	for ; p*10 <= num; p *= 10 {
		curlength++
	}

	// part that prints spaces so that matrix numbers arr printed correctly
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

// function scans width and height of map. Retuns scanned values
func readWidthAndHeight() (int, int) {
	var width, height int

	if _, err := fmt.Scanf("%d %d", &height, &width); err != nil {
		return -1, -1
	}

	// make needed to implement Map numbers on x and y-axises.
	return height, width
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
			if full_map[i][j] = int(cell - '0'); full_map[i][j] < 0 || full_map[i][j] > 3 {
				return nil
			}

		}
		if _, err := fmt.Scanf("%c", &cell); err != nil {
			return nil
		}
	}

	return full_map
}

// function to read change symbols
func readFormat() {
	temp := make([]rune, 3)

	for i := 0; i < 3; i++ {
		if _, err := fmt.Scanf("%c", &temp[i]); temp[i] == '\n' || temp[i] < 32 || temp[i] > 126 || err != nil {
			colorizeFG(0)
			printString("\nInvalid input for change symbols\n")
			printString("Default values (X>@) will be used\n\n")
			colorizeFG(1)
			return
		}
	}
	wall = temp[0]
	player = temp[1]
	award = temp[2]
}

// function to provide clear instructions for input format.
func greetingsMsg() {
	printString("\nHello this is program that prints a map from a given input.\n")

	colorizeBG(0)
	printString("Only ASCII characters are allowed to use.")
	colorizeBG(1)

	printString("\n\nThe map consists of the following characters:\n")
	colorizeFG(0)
	printString("2 - a player.\n")
	printString("0 - a wall, where the player is not allowed to move.\n")
	printString("1 - a free cell, where the player is allowed to step in.\n")
	printString("3 - an award.")
	colorizeFG(1)
	printString("\n\n============================================================================\n\n")
	colorizeBG(2)
	printString("The input is of the following format: ")
	colorizeBG(1)
	printString("\nh w, where h is height of map and w is width of map.\n\n")
	printString("After that come h lines each with w number of characters.\n")

	printString("\nFinally enter change symbols for display of cells.\n")
	colorizeFG(3)
	printString("in this order: \n1. Wall\n2. Player\n3. Award\n")
	colorizeBG(1)
	printString("\n\n============================================================================\n\n")
	colorizeFG(1)
	colorizeBG(1)
}
