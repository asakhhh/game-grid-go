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

// printCell function prints a single cell based on its value
func printCell(value int, x int, y int) {
	if value == 0 {
		fmt.Print(RED + "X" + RESET) // Wall
	} else if value == 1 {
		if x%3 == 0 {
			fmt.Print("_") // Horizontal separator
		} else {
			fmt.Print(" ") // Empty space
		}
	} else if value == 2 {
		if x == 2 && y == 4 {
			fmt.Print(BLUE + string(player) + RESET) // Player
		} else if x == 0 {
			fmt.Print("_") // Horizontal separator
		} else {
			fmt.Print(" ") // Empty space
		}
	} else if value == 3 {
		if x == 2 && y == 4 {
			fmt.Print(YELLOW + string(award) + RESET) // Award
		} else if x == 0 {
			fmt.Print("_") // Horizontal separator
		} else {
			fmt.Print(" ") // Empty space
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
			fmt.Printf("%2d", (h+1)/3) // Print vertical coordinate every 3rd row
		} else {
			fmt.Print("  ") // Print space for alignment
		}

		// Loop to iterate over every column
		// Multiplied by 8 because every column has 7 runes inside and borders
		if h == 0 {
			for w := 0; w < (8*width)+1; w++ {
				if (w+3)%7 == 0 && w != 0 {
					fmt.Print(string(rune(64 + (w+3)/7))) // Print horizontal coordinates
				} else {
					fmt.Print(" ") // Print space for alignment
				}
			}
			fmt.Print("\n") // Move to the next line after printing coordinates
		}

		for w := 0; w < (8*width)+1; w++ {
			if w == 0 && h != 0 || w%8 == 0 && h != 0 {
				fmt.Print("|") // Vertical border
			} else if h == 0 {
				if w == 0 {
					fmt.Print("  ")
				}
				fmt.Print("_") // Horizontal border
			} else if h%3 != 0 {
				printCell(value[h/3][w/8], h%3, w%8) // Print cell content
			} else {
				printCell(value[h/3-1][w/8], h%3, w%8) // Print cell content
			}
		}
		fmt.Print("\n") // Move to the next line after printing a row
	}
}
