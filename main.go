package main

import "fmt"

var (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	WHITE  = "\033[97m"
	BLUE   = "\033[34m"
	YELLOW = "\033[33m"
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

	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			fmt.Scanf("%d", &full_map[i][j])
		}
	}

	printMap(HEIGHT, WIDTH, full_map)
}

func printCell(value int, x int, y int) {
	if value == 0 {
		fmt.Print(RED + "X" + RESET)
	} else if value == 1 {
		if x%3 == 0 {
			fmt.Print("_")
		} else {
			fmt.Print(" ")
		}
	} else if value == 2 {
		fmt.Print(BLUE + ">" + RESET)
	} else if value == 3 {
		fmt.Print(YELLOW + "@" + RESET)
	}
}

func printMap(height int, width int, value [][]int) {
	// Loop to iterate over every row.
	// Multiplied by 3 because every row has 3 rows inside.
	// And additional 1 to print first line
	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 && h != 0 {
			fmt.Printf("%2d", (h+1)/3)
		} else {
			fmt.Print("  ")
		}

		// Loop to iterate over every column
		if h == 0 {
			for w := 0; w < (7*width)+width+1; w++ {
				if (w+3)%7 == 0 && w != 0 {
					fmt.Print(string(rune(64 + (w+3)/7)))
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}

		for w := 0; w < (8*width)+1; w++ {
			if w == 0 && h != 0 || w%8 == 0 && h != 0 {
				fmt.Print("|")
			} else if h == 0 {
				if w == 0 {
					fmt.Print("  ")
				}
				fmt.Print("_")
			} else if h%3 != 0 {
				printCell(value[h/3][w/8], h, w)
			} else {
				printCell(value[h/3-1][w/8], h, w)
			}
		}
		fmt.Print("\n")
	}
}
